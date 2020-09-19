package zzeutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

//返回指定路径的文件大小，单位为 Byte
func GetFileSize(filename string) uint64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			result = 0
		} else {
			result = f.Size()
		}
		return nil
	})
	return uint64(result)
}

func BytesToMB(size uint64) float64 {
	mbSize, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(size)/1024/1024), 64)
	return mbSize
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
