package cmd

import (
	"github.com/spf13/cobra"
	"go-simple/pkg/console"
	"go-simple/pkg/helpers"
)

var Key = &cobra.Command{
	Use: "key",
	Short: "Generate App Key, will print the generated Key",
	Run: runKeyGenerate,
	Args:  cobra.NoArgs, // Parameter transfer is not allowed
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("---")
	console.Success("App Key:")
	console.Success("-" + helpers.RandomString(32) + "-")
	console.Success("---")
	console.Warning("please go to .env file to change the APP_KEY option")
}