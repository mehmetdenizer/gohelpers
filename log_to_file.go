package gohelpers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

// LogError centralizes logging for errors with a specific context
func LogError(msg string, err error) {
	var message string
	if err != nil {
		message = fmt.Sprintf("%s: %v", msg, err)
	} else {
		message = msg
	}

	logToFile(message, "error")

	log.Printf(message)
}

// LogInfo centralizes logging for infos with a specific message
func LogInfo(message string) {
	logToFile(message, "info")
}

// LogWarning centralizes logging for warnings with a specific message
func LogWarning(message string) {
	logToFile(message, "warning")
}

func logToFile(message, kind string) {
	logPath := fmt.Sprintf("storage/logs/%s", kind)
	createLogPath(logPath)

	fileName := logFileName(logPath, kind)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("goHelpers: Log file could not be opened: %v", err)
	}
	defer closeFile(file)

	logrus.SetOutput(file)
	logByKind(kind, message)

}

func createLogPath(path string) {
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatalf("goHelpers: Path could not be created: %v", err)
	}
}

func logFileName(path string, kind string) string {
	now := time.Now()
	return fmt.Sprintf("%s/%d-%02d_%s.log", path, now.Year(), int(now.Month()), kind)
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Fatal("goHelpers: An error occurred while closing the file:", err)
	}
}

func logByKind(kind, message string) {
	switch kind {
	case "info":
		logrus.Info(message)
	case "warning":
		logrus.Warn(message)
	case "error":
		logrus.Error(message)
	default:
		logrus.Errorf("goHelpers: Invalid log kind: %s. Message: %s", kind, message)
	}
}
