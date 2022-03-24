package logger

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/termenv"
)

var profile = termenv.TrueColor
var redForeground = profile.FromColor(colorful.LinearRgb(1, 0.5, 0.5))

//245, 166, 132
var orangeFg = profile.FromColor(colorful.LinearRgb(245.0/255, 166.0/255, 132.0/255))
var redBg = profile.FromColor(colorful.LinearRgb(224.0/255, 20.0/255, 68.0/255))
var orange = profile.FromColor(colorful.LinearRgb(0.9, 0.3, 0.0))
var white = profile.FromColor(termenv.ConvertToRGB(termenv.ANSIBrightWhite))

func Warning(text string) {
	s := termenv.String(" Warning ")
	//#ffa442
	s = s.
		Background(orange).
		Foreground(white)
	t := termenv.String(" " + text + " ").Foreground(orangeFg)
	fmt.Println(s.String() + " " + t.String())
}

func WarningNoTag(text string) {
	t := termenv.String(text).Foreground(orangeFg)
	fmt.Println(t.String())
}

func Error(kind, text string, paragraph ...interface{}) {
	s := termenv.String(" Error " + kind + " ")
	//#ffa442
	s = s.
		Background(redBg).
		Foreground(white)
	t := termenv.String(" " + text + " ").Foreground(redForeground)
	fmt.Println(s.String() + " " + t.String())
	fmt.Println(paragraph...)
}

func Success(text string) {
	s := termenv.String(text)
	s = s.
		Foreground(termenv.ANSIGreen)
	fmt.Println(s)
}
