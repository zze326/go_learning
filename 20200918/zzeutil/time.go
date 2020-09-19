package zzeutil

import "time"

func NowStr(format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05.000"
	}
	return time.Now().Format(format)
}
