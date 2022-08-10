package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-simple/app/middlewares"
	"go-simple/routes"
	"net/http"
	"strings"
)

func SetupRoute(router *gin.Engine) {
	registerGlobalMiddleWare(router)

	setupNotFoundHandle(router)

	routes.Initialize()
}

// Global Middleware
func registerGlobalMiddleWare(router *gin.Engine)  {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setupNotFoundHandle(router *gin.Engine)  {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":http.StatusNotFound,
				"error_message":"路由未定义，请确认 url 和请求方法是否正确",
			})
		}
	})
}


