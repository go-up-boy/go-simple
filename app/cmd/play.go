package cmd

import (
	"github.com/spf13/cobra"
)

var Play = &cobra.Command{
	Use: "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run: runPay,
}

func runPay(cmd *cobra.Command, args []string) {

}