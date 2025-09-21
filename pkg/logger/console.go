package logger

import "fmt"

type ConsoleLogger struct{}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (l *ConsoleLogger) Info(msg string, args ...any) {
	fmt.Printf("INFO: "+msg+"\n", args...)
}

func (l *ConsoleLogger) Error(msg string, args ...any) {
	fmt.Printf("ERROR: "+msg+"\n", args...)
}
