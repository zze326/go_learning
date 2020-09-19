package zlogger

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type LogLevel int8

const (
	Debug LogLevel = iota
	Trace LogLevel = iota
	Info  LogLevel = iota
	Warn  LogLevel = iota
	Error LogLevel = iota
	Off   LogLevel = iota
)

func parseLogLevel(levelStr string) (logLevel LogLevel, err error) {
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		logLevel = Debug
	case "trace":
		logLevel = Trace
	case "info":
		logLevel = Info
	case "warn":
		logLevel = Warn
	case "error":
		logLevel = Error
	case "off":
		logLevel = Off
	case "all":
		logLevel = Debug
	default:
		err = errors.New(fmt.Sprintf("parse loglevel failed, can not support this loglevel: %v", levelStr))
	}
	return
}

type Logger interface {
	log(logLevel LogLevel, format string, params ...interface{})
	Debug(format string, params ...interface{})
	Trace(format string, params ...interface{})
	Info(format string, params ...interface{})
	Warn(format string, params ...interface{})
	Error(format string, params ...interface{})
}

func getRuntimeInfo() (filePath, funcName string, line int) {
	pc, filePath, line, _ := runtime.Caller(3)
	f := runtime.FuncForPC(pc)
	funcName = f.Name()
	return
}

func formatLogMsg(levelStr, funcName, filePath, nowStr, msg string, line int) string {
	msg = fmt.Sprintf("[%s] [%s] [%s:%s:%d] [%s]\n", nowStr, levelStr, filePath, funcName, line, msg)
	return msg
}
