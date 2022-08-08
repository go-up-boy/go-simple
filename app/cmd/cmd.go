package cmd

import (
	"github.com/spf13/cobra"
	"go-simple/pkg/helpers"
	"os"
)

// Env 存储全局选项 --env 的值
var Env string

func RegisterGlobalFlags(rootCmd *cobra.Command)  {
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .env .file, example:--env=testing will use .env.testing file")
}

func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArt := helpers.FirstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firstArt != "-h" && firstArt != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}