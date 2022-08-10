package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-simple/pkg/console"
	"go-simple/pkg/helpers"
	"os"
)

var CmdMakeAPIController = &cobra.Command{
	Use: "apicontroller",
	Short: "Create api controller，exmaple: make apicontroller api/v1/user",
	Run: runMakeAPIController,
	Args: cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	array := helpers.RemoveEmptyToArray(args[0])
	if len(array) < 1 {
		console.Error("please input path/controllerName")
	}
	dir := helpers.RemoveEmptyToString(array[:len(array) - 1])
	model := makeModelFromString(array[len(array) - 1])
	filePath := fmt.Sprintf("app/http/controllers/%s/", dir + "/" + model.PackageName)
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		console.Error("apicontroller filePath create fail " + err.Error())
	}
	createFileFromStub(filePath + model.TableName + "_controller.go", "apicontroller", model)
}
