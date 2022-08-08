package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-simple/app/modules/user_module/user"
	"go-simple/pkg/jwt"
	"go-simple/pkg/response"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(c)
		if err != nil {
			response.Unauthorized(c, err.Error())
			return
		}
		userModel := user.Get(claims.UserId)
		if userModel.ID == 0 {
			response.Unauthorized(c, "用户不存在")
			return
		}
		c.Set("current_user_id", userModel.ID)
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}