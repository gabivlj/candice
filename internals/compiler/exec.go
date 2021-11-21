package compiler

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
)

func GenerateExecutable(writer io.WriterTo, path string) error {
	_ = os.Remove(".intermediate_output.ll")
	fd, _ := os.Create(".intermediate_output.ll")
	_, _ = writer.WriteTo(fd)
	cmd := exec.Command("clang", ".intermediate_output.ll", "-o", path, "pr.o")
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stdout
	err := cmd.Run()
	_ = os.Remove(".intermediate_output.ll")
	if err != nil {
		return errors.New("error compiling with clang:\n" + stdout.String())
	}
	return nil
}
