package compiler

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	l "tinygo.org/x/go-llvm"
)

func GenerateExecutable(writer io.WriterTo, path string) error {
	_ = os.Remove(".intermediate_output.ll")
	fd, _ := os.Create(".intermediate_output.ll")
	_, _ = writer.WriteTo(fd)
	cmd := exec.Command("clang", ".intermediate_output.ll", "-o", path)
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stdout
	err := cmd.Run()
	//_ = os.Remove(".intermediate_output.ll")
	if err != nil {
		return errors.New("error compiling with clang:\n" + stdout.String())
	}
	return nil
}

func GenerateObjectLLVM(writer io.WriterTo, path string) error {
	_ = os.Remove(".intermediate_output.ll")
	intermediateOutputFd, err := os.Create(".intermediate_output.ll")
	if err != nil {
		return err
	}
	_, err = writer.WriteTo(intermediateOutputFd)
	ctx := l.GlobalContext()
	if err != nil {
		return err
	}
	l.InitializeAllTargetInfos()
	l.InitializeAllTargets()
	l.InitializeAllTargetMCs()
	l.InitializeAllAsmParsers()
	l.InitializeAllAsmPrinters()

	mem, err := l.NewMemoryBufferFromFile(".intermediate_output.ll")
	if err != nil {
		return err
	}
	module, err := ctx.ParseIR(mem)
	if err != nil {
		return err
	}
	tripleTarget := l.DefaultTargetTriple()
	target, err := l.GetTargetFromTriple(tripleTarget)
	if err != nil {
		return err
	}
	model := l.RelocDefault
	targetMachine := target.CreateTargetMachine(tripleTarget, "generic", "", l.CodeGenLevelAggressive, model, l.CodeModelDefault)
	module.SetTarget(tripleTarget)
	module.SetDataLayout(targetMachine.CreateTargetData().String())
	passManager := l.NewPassManager()
	targetMachine.AddAnalysisPasses(passManager)
	memBuffer, err := targetMachine.EmitToMemoryBuffer(module, l.ObjectFile)
	if err != nil {
		return err
	}
	fdOutput, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = fdOutput.Write(memBuffer.Bytes())
	if err != nil {
		return err
	}

	return err
}
