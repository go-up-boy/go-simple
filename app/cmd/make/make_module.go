package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-simple/pkg/console"
	"os"
)

var CmdMakeModule = &cobra.Command{
	Use: "module",
	Short: "Create a set of modules",
	Run: runModule,
	Args: cobra.ExactArgs(2),
}

func runModule(cmd *cobra.Command, args []string) {
	moduleName := args[0]
	controllerPath := args[1]
	model := makeModelFromString(moduleName)
	createModuleContents(model.PackageName)
	path := fmt.Sprintf("app/modules/%s_module/%s_logics/logic.go", model.PackageName, model.PackageName)
	createFileFromStub(path, "logic", model)
	runMakeRequest(cmd, []string{moduleName})
	runMakeAPIController(cmd, []string{controllerPath + "/" + moduleName})
	runMakeModel(cmd, []string{moduleName})
	runMakeRoute(cmd, []string{controllerPath + "/" + moduleName})
}

func createModuleContents(moduleName string) error {
	moduleContents := []string{
		moduleName + "/",
		moduleName + "_logics/",
		moduleName + "_services/",
		moduleName + "_utils/",
		moduleName + "_policies/",
	}
	for _,dir := range moduleContents{
		path := fmt.Sprintf("app/modules/%s_module/%s/", moduleName, dir)
		err := os.MkdirAll(path, os.ModePerm)
		console.ExitIf(err)
	}
	return nil
}