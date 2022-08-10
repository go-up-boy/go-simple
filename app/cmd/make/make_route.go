package make

import (
	"github.com/spf13/cobra"
	"go-simple/pkg/console"
	"go-simple/pkg/helpers"
)

var CmdMakeRoute = &cobra.Command{
	Use: "route",
	Short: "create controller route",
	Run: runMakeRoute,
	Args: cobra.ExactArgs(1),
}

func runMakeRoute(cmd *cobra.Command, args []string) {
	array := helpers.RemoveEmptyToArray(args[0])
	if len(array) < 2 {
		console.Error("The path must be greater than two levels")
	}
	dir := helpers.RemoveEmptyToString(array[:len(array) - 1])
	model := makeModelFromString(array[len(array) - 1])
	model.ControllerPath = dir
	createFileFromStub("routes/" + model.TableName + ".go", "route", model)
}
