package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/redis"
	"go-simple/pkg/config"
	"go-simple/pkg/logger"
	redispkg "go-simple/pkg/redis"
	"strings"
)

func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

func GetRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}

func CheckRate(c *gin.Context, key string, formatted string) (limiter.Context, error) {
	var context limiter.Context
	rate, err := limiter.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context, err
	}
	store, err := redis.NewStoreWithOptions(redispkg.Redis.Client, limiter.StoreOptions{
		Prefix: config.GetString("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	limiterObj := limiter.New(store, rate)
	if c.GetBool("limiter-once") {
		return limiterObj.Peek(c, key)
	} else {
		c.Set("limiter-once", true)
		return limiterObj.Get(c, key)
	}
}

// routeToKeyString 辅助方法，将 URL 中的 / 格式为 -
func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
