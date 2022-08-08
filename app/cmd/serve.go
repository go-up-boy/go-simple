package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-simple/bootstrap"
	"go-simple/pkg/config"
	"go-simple/pkg/console"
	"go-simple/pkg/logger"
)

var Serve = &cobra.Command{
	Use: "serve",
	Short: "Start web server",
	Run: runWeb,
	Args: cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string)  {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	bootstrap.SetupRoute(router)
	// 运行服务器
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}