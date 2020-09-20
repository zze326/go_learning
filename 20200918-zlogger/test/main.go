package main

import (
	"testProject/20200918/zlogger"
)

func main() {
	//logger := zlogger.CreateConsoleLogger("info")
	logger := zlogger.CreateFileLogger("DEBUG", "./", "zlog", 1024*1024)
	for i := 0; i < 1000; i++ {
		logger.Debug("hello zze %d", i)
		logger.Trace("hello zze %d", i)
		logger.Info("hello zze %d", i)
		logger.Warn("hello zze %d", i)
		logger.Error("hello zze %d", i)
	}
}
