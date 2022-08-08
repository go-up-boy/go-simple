package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-simple/app"
	"go-simple/pkg/limiter"
	"go-simple/pkg/logger"
	"go-simple/pkg/response"
	"net/http"
)

func LimitIP(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}
	return func(c *gin.Context) {
		key := limiter.GetKeyIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func LimitPreRoute(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}
	return func(c *gin.Context) {
		// 针对单个路由，增加访问次数
		c.Set("limiter-once", false)
		// 对ip + 路由 限流
		key := limiter.GetRouteWithIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func limitHandler(c *gin.Context, key string, formatted string) bool {
	rate, err := limiter.CheckRate(c, key, formatted)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c, err.Error())
		return false
	}
	// ---- 设置标头信息-----
	// X-RateLimit-Limit :10000 最大访问次数
	// X-RateLimit-Remaining :9993 剩余的访问次数
	// X-RateLimit-Reset :1513784506 到该时间点，访问次数会重置为 X-RateLimit-Limit
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))
	if rate.Reached {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "接口请求太频繁",
		})
		return false
	}
	return true
}