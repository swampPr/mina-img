// Package logger provides logger  ->  Provides a logger to os.Stdout
package logger

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	green  = color.New(color.FgGreen).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
)

type Logger struct{}

func (l *Logger) Info(message string) {
	fmt.Println(green(message))
}

func (l *Logger) Error(message string) {
	fmt.Println(red(message))
}

func (l *Logger) Warn(message string) {
	fmt.Println(yellow(message))
}
