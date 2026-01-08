package main

import (
	"fmt"
)

type Logger struct{}

func (l *Logger) Info(msg string) {
	fmt.Println("ℹ️ ", msg)
}

func (l *Logger) Success(msg string) {
	fmt.Println("✅", msg)
}

func (l *Logger) Error(msg string) {
	fmt.Println("❌", msg)
}
