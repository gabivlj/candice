package terminal

import (
	"os/exec"
)

func ProgramExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func ClangExists() bool {
	return ProgramExists("clang")
}
