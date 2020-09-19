package zzeutil

import (
	"fmt"
)

type FontColor int8

const (
	//黑色
	Black FontColor = iota
	//红色
	Red FontColor = iota
	//绿色
	Green FontColor = iota
	//黄色
	Yellow FontColor = iota
	//蓝色
	Blue FontColor = iota
	//紫色
	Purple FontColor = iota
	//深绿色
	DarkGreen FontColor = iota
	//白色
	White FontColor = iota
	//白底黑字
	WhiteReverse FontColor = iota
	//深绿底红字
	DarkGreenReverse FontColor = iota
	//紫底绿字
	PurpleReverse FontColor = iota
	//蓝底黄字
	BlueReverse FontColor = iota
	//黄底蓝字
	YellowReverse FontColor = iota
	//绿底紫字
	GreenReverse FontColor = iota
	//红底深绿字
	RedReverse FontColor = iota
	//黑底白字
	BlackReverse FontColor = iota
)

func ColorPrint(fontColor FontColor, format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)

	colorFormat1 := "\x1b[%dm%%s\x1b[0m"
	colorFormat2 := "\x1b[%d;%dm%%s\x1b[0m"
	switch fontColor {
	case Black:
		format = fmt.Sprintf(colorFormat1, 30)
	case Red:
		format = fmt.Sprintf(colorFormat1, 31)
	case Green:
		format = fmt.Sprintf(colorFormat1, 32)
	case Yellow:
		format = fmt.Sprintf(colorFormat1, 33)
	case Blue:
		format = fmt.Sprintf(colorFormat1, 34)
	case Purple:
		format = fmt.Sprintf(colorFormat1, 35)
	case DarkGreen:
		format = fmt.Sprintf(colorFormat1, 36)
	case White:
		format = fmt.Sprintf(colorFormat1, 37)
	case WhiteReverse:
		format = fmt.Sprintf(colorFormat2, 47, 30)
	case DarkGreenReverse:
		format = fmt.Sprintf(colorFormat2, 46, 31)
	case PurpleReverse:
		format = fmt.Sprintf(colorFormat2, 45, 32)
	case BlueReverse:
		format = fmt.Sprintf(colorFormat2, 44, 33)
	case YellowReverse:
		format = fmt.Sprintf(colorFormat2, 43, 34)
	case GreenReverse:
		format = fmt.Sprintf(colorFormat2, 42, 35)
	case RedReverse:
		format = fmt.Sprintf(colorFormat2, 41, 36)
	case BlackReverse:
		format = fmt.Sprintf(colorFormat2, 40, 37)
	}
	fmt.Printf(format, msg)
}

//返回一个创造打印指定颜色内容到终端的闭包
func CreateColorPrinter(color FontColor) func(format string, params ...interface{}) {
	return func(format string, params ...interface{}) {
		ColorPrint(color, format, params...)
	}
}
