package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-simple/app/cmd"
	cmdMake "go-simple/app/cmd/make"
	"go-simple/bootstrap"
	btsConfig "go-simple/config"
	"go-simple/globals"
	"go-simple/pkg/config"
	"go-simple/pkg/console"
	"os"
)

func init() {
	// 加载应用程序配置文件
	btsConfig.Initialize()
}

func main()  {
	var rootCmd = &cobra.Command{
		Use: config.Get("app.name"),
		Short: "A simple forum project",
		Long: `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		PersistentPreRun: func(command *cobra.Command, args []string) {
			config.InitConfig(cmd.Env)
			bootstrap.SetupLogger()
			bootstrap.SetupDB()
			bootstrap.SetupRedis()
			globals.GlobalLazyInit()
			bootstrap.SetupRoute(globals.GlobalService.R)
		},
	}
	rootCmd.AddCommand(
			cmd.Serve,
			cmd.Key,
			cmd.Play,
			cmdMake.Make,
			cmd.Migrate,
		)
	cmd.RegisterGlobalFlags(rootCmd)
	cmd.RegisterDefaultCmd(rootCmd, cmd.Serve)
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v : %s", os.Args, err.Error()))
	}
}
