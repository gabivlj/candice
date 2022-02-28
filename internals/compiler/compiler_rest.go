//go:build windows || linux || amd64
// +build windows linux amd64

package compiler

import (
	"github.com/gabivlj/candice/pkg/logger"
	"io"
	"os"
)

func GenerateObjectLLVM(writer io.WriterTo, path string, _ bool) (string, error) {
	logger.Warning("You are using a Clang and LLVM dependant build of candice, consider contributing by building candice on your platform!")
	_ = os.Remove(path)
	intermediateOutputFd, err := os.Create(path + ".ll")
	if err != nil {
		return "", err
	}
	_, err = writer.WriteTo(intermediateOutputFd)
	return path + ".ll", err
}
