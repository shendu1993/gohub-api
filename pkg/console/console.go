package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

// Success 打印一条成功消息，绿色输出
func Success(msg string) {
	colorOut(msg, "green")
}

// Error 打印一条报错消息，红色输出
func Error(msg string) {
	colorOut(msg, "red")
}

// Warning 打印一条提示消息，黄色输出
func Warning(msg string) {
	colorOut(msg, "yellow")
}

// Exit 打印一条报错消息，并退出os.Exit(1)
func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

//colorOut 内部使用 设置颜色高亮
func colorOut(message, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(message, color))
}