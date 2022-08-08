package console

import (
	"fmt"
	"github.com/mgutz/ansi"
	"go-simple/app"
	"os"
)

func Success(msg string) {
	colorOut(msg, "green")
}

func Error(msg string) {
	colorOut(msg, "red")
}

func Warning(msg string) {
	colorOut(msg, "yellow")
}

func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

// ExitIf 语法糖，自带 err != nil 判断
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

func colorOut(message, color string)  {
	if app.IsWindows() {
		fmt.Println(message)
	} else {
		fmt.Fprintln(os.Stdout, ansi.Color(message, color))
	}
}