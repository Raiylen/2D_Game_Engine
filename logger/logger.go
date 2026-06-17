package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	reset   string = "\033[0m"
	red     string = "\033[31m"
	bRed    string = "\033[1;31m"
	yellow  string = "\033[33m"
	bYellow string = "\033[1;33m"
	green   string = "\033[32m"
	bGreen  string = "\033[1;32m"
	cyan    string = "\033[36m"
)

type logType int

const (
	logInfo logType = iota
	logWarning
	logError
)

type logEntry struct {
	logType logType
	message string
}

var logs []logEntry

func Info(info string) {
	timeNow := time.Now().Format(time.DateTime)
	message := fmt.Sprintf("%sINFO: | %v — %s%s\n", bGreen, timeNow, info, reset)
	fmt.Fprint(os.Stderr, message)

	entry := logEntry{logType: logInfo, message: message}
	logs = append(logs, entry)
}

func Warn(warning string) {
	timeNow := time.Now().Format(time.DateTime)
	message := fmt.Sprintf("%sWARN: | %v — %s%s\n", bYellow, timeNow, warning, reset)
	fmt.Fprint(os.Stderr, message)

	entry := logEntry{logType: logInfo, message: message}
	logs = append(logs, entry)
}

func Err(error error) {
	timeNow := time.Now().Format(time.DateTime)
	message := fmt.Sprintf("%sERR: | %v — %s%s\n", bRed, timeNow, error, reset)
	fmt.Fprintf(os.Stderr, message)

	entry := logEntry{logType: logInfo, message: message}
	logs = append(logs, entry)
}
