package factory

import (
	"fmt"
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type ConsoleLogger struct{}

func (l *ConsoleLogger) Info(msg string) {
	println("INFO:", msg)
}

func (l *ConsoleLogger) Error(msg string) {
	println("ERROR:", msg)
}

type FileLogger struct {
	f *os.File
}

func NewFileLogger(path string) (*FileLogger, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileLogger{f}, nil
}

func (l *FileLogger) Info(msg string) {
	l.f.WriteString("INFO: " + msg + "\n")
}

func (l *FileLogger) Error(msg string) {
	l.f.WriteString("ERROR: " + msg + "\n")
}

func NewLogger(loggerType string, path string) (Logger, error) {
	switch loggerType {
	case "console":
		return &ConsoleLogger{}, nil
	case "file":
		return NewFileLogger(path)
	default:
		return nil, fmt.Errorf("unsupported logger type: %s", loggerType)
	}
}

func RunLoggerDemo() {
	logger, err := NewLogger("console", "")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	logger.Info("This is an info message")
	logger.Error("This is an error message")

	fileLogger, err := NewLogger("file", "app.log")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fileLogger.Info("This is an info message in file")
	fileLogger.Error("This is an error message in file")
}
