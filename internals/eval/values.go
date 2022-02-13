package eval

import (
	"runtime"

	"github.com/gabivlj/candice/internals/token"
)

type Value interface {
	IsTruthy() bool
	constantValue()
}

type Integer struct {
	Value int64
}

func (i *Integer) constantValue() {}

func (i *Integer) IsTruthy() bool {
	return i.Value != 0
}

type Error struct {
	Message string
	Token   token.Token
}

func (e *Error) constantValue() {}

func (e *Error) IsTruthy() bool {
	return false
}

func boolToInteger(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

var constants = map[string]Value{
	"WINDOWS": &Integer{
		Value: boolToInteger(runtime.GOOS == "windows"),
	},
	"MACOS": &Integer{
		Value: boolToInteger(runtime.GOOS == "darwin"),
	},
	"LINUX": &Integer{
		Value: boolToInteger(runtime.GOOS == "linux"),
	},
	"X64": &Integer{
		Value: boolToInteger(runtime.GOARCH == "amd64"),
	},
	"ARM64": &Integer{
		Value: boolToInteger(runtime.GOARCH == "arm64"),
	},
	"ARM": &Integer{
		Value: boolToInteger(runtime.GOARCH == "arm"),
	},
	"ARM64BE": &Integer{
		Value: boolToInteger(runtime.GOARCH == "arm64be"),
	},
	"ARMBE": &Integer{
		Value: boolToInteger(runtime.GOARCH == "armbe"),
	},
	"386": &Integer{
		Value: boolToInteger(runtime.GOARCH == "386"),
	},
}
