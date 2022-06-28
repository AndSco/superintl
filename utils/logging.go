package utils

import (
	"fmt"
)

type Color string

const (
	colorReset   Color = "\033[0m"
	colorRed     Color = "\033[31m"
	colorGreen   Color = "\033[32m"
	colorYellow  Color = "\033[33m"
	colorBlue    Color = "\033[0;36m"
	colorMagenta Color = "\u001b[35m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color) + message + string(colorReset))
}

func Alert(message string) {
	AddLine(1)
	colorize(colorRed, message)
	AddLine(1)
}

func Warn(message string) {
	colorize(colorYellow, message)
}

func Warn2(message string) {
	AddLine(1)
	colorize(colorMagenta, message)
	AddLine(1)
}

func Inform(message string) {
	AddLine(1)
	colorize(colorBlue, message)
	AddLine(1)
}

func Success(message string) {
	// AddLine(1)
	colorize(colorGreen, message)
	// AddLine(1)
}

func AddLine(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("")
	}
}
