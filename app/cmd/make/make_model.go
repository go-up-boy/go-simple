package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdMakeModel = &cobra.Command{
	Use: "model",
	Short: "Crate model file, example: make model user",
	Run: runMakeModel,
	Args: cobra.ExactArgs(1),
}

func runMakeModel(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	// 模型目录验证
	modelDir := fmt.Sprintf("app/modules/%s_module/%s/", model.PackageName, model.PackageName)
	os.MkdirAll(modelDir, os.ModePerm)
	// 替换变量
	createFileFromStub(modelDir + model.PackageName + "_model.go", "model/model", model)
	createFileFromStub(modelDir + model.PackageName + "_util.go", "model/model_util", model)
	createFileFromStub(modelDir + model.PackageName + "_hooks.go", "model/model_hooks", model)
}
