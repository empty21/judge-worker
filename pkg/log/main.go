package log

import "fmt"

func Error(message string, args ...any) {
	log("error", fmt.Sprintf(message, args...))
}

func Info(message string, args ...any) {
	log("info", fmt.Sprintf(message, args...))
}

func Debug(message string, args ...any) {
	log("debug", fmt.Sprintf(message, args...))
}
