package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go-simple/app"
	"go-simple/pkg/config"
	"go-simple/pkg/logger"
	"strings"
	"time"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

type JWT struct {
	SignKey []byte
	MaxRefresh time.Duration
}

type CustomClaims struct {
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	ExpireAtTime int64 `json:"expire_time"`

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey: []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

func (j *JWT) ParserToken(c *gin.Context) (*CustomClaims, error) {
	tokenString, parseErr := j.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}
	token, err := j.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if ok {
			if validationErr.Errors == jwt.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwt.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

func (j *JWT) RefreshToken(c *gin.Context) (string, error) {
	tokenString, parseErr := j.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}
	token, err := j.parseTokenString(tokenString)
	// 3. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors != jwt.ValidationErrorExpired {
			return "", err
		}
	}
	// 解析CustomClaims 数据
	claims := token.Claims.(*CustomClaims)
	x := app.TimeNowInTimezone().Add(-j.MaxRefresh).Unix()
	if claims.IssuedAt > x {
		claims.StandardClaims.ExpiresAt = j.expireAtTime()
		return j.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

func (j *JWT) IssueToken(userID string, userName string) string {
	expireAtTime := j.expireAtTime()
	claims := CustomClaims{
		userID,
		userName,
		expireAtTime,
		jwt.StandardClaims{
			NotBefore: app.TimeNowInTimezone().Unix(), // 签名生效时间
			IssuedAt:  app.TimeNowInTimezone().Unix(), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: expireAtTime,                   // 签名过期时间
			Issuer:    config.GetString("app.name"),   // 签名颁发者
		},
	}
	token, err := j.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

func (j *JWT) createToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}

// expireAtTime 过期时间
func (j *JWT) expireAtTime() int64 {
	timeNow := app.TimeNowInTimezone()

	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}

	expire := time.Duration(expireTime) * time.Minute
	return timeNow.Add(expire).Unix()
}

func (j JWT) parseTokenString(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})
}

func (j *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 && parts[0] != "Bearer" {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}