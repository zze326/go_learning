package zlogger

import (
	"fmt"
	"testProject/utils"
	"time"
)

//ConsoleLogger 控制台 Logger
type ConsoleLogger struct {
	logLevel LogLevel
}

//CreateConsoleLogger ConsoleLogger 构造函数
func CreateConsoleLogger(levelStr string) (logger *ConsoleLogger) {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &ConsoleLogger{
		logLevel: logLevel,
	}
}

//log 公用日志输出函数
func (c *ConsoleLogger) log(logLevel LogLevel, format string, params ...interface{}) {
	if c.logLevel > logLevel {
		return
	}
	msg := fmt.Sprintf(format, params...)
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	var levelStr string
	var fontColor utils.FontColor
	switch logLevel {
	case Debug:
		levelStr = "DEBUG"
		fontColor = utils.Black
	case Trace:
		levelStr = "TRACE"
		fontColor = utils.White
	case Info:
		levelStr = "INFO"
		fontColor = utils.Blue
	case Warn:
		levelStr = "WARN"
		fontColor = utils.Yellow
	case Error:
		levelStr = "ERROR"
		fontColor = utils.Red
	}
	filePath, funcName, line := getRuntimeInfo()
	//fileName := filepath.Base(filePath)
	msg = formatLogMsg(levelStr, funcName, filePath, nowStr, msg, line)
	utils.ColorPrint(fontColor, msg)
}

//Debug ...
func (c *ConsoleLogger) Debug(format string, params ...interface{}) {
	c.log(Debug, format, params...)
}

//Trace ...
func (c *ConsoleLogger) Trace(format string, params ...interface{}) {
	c.log(Trace, format, params...)
}

//Info ...
func (c *ConsoleLogger) Info(format string, params ...interface{}) {
	c.log(Info, format, params...)
}

//Warn ...
func (c *ConsoleLogger) Warn(format string, params ...interface{}) {
	c.log(Warn, format, params...)
}

//Error ...
func (c *ConsoleLogger) Error(format string, params ...interface{}) {
	c.log(Error, format, params...)
}
