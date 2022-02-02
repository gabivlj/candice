//go:build darwin && arm64

package compiler

import (
	"io"
	"os"

	l "tinygo.org/x/go-llvm"
)

func GenerateObjectLLVM(writer io.WriterTo, path string, optimized bool) (string, error) {
	defer func() {
		_ = os.Remove(".intermediate_output.ll")
	}()
	intermediateOutputFd, err := os.Create(".intermediate_output.ll")
	if err != nil {
		return "", err
	}
	_, err = writer.WriteTo(intermediateOutputFd)
	ctx := l.GlobalContext()
	if err != nil {
		return "", err
	}

	l.InitializeAllTargetInfos()
	l.InitializeAllTargets()
	l.InitializeAllTargetMCs()
	l.InitializeAllAsmParsers()
	l.InitializeAllAsmPrinters()

	mem, err := l.NewMemoryBufferFromFile(".intermediate_output.ll")
	if err != nil {
		return "", err
	}
	module, err := ctx.ParseIR(mem)
	if err != nil {
		return "", err
	}
	tripleTarget := l.DefaultTargetTriple()
	target, err := l.GetTargetFromTriple(tripleTarget)
	if err != nil {
		return "", err
	}
	model := l.RelocDefault
	targetMachine := target.CreateTargetMachine(tripleTarget, "generic", "", l.CodeGenLevelAggressive, model, l.CodeModelSmall)
	module.SetTarget(tripleTarget)
	module.SetDataLayout(targetMachine.CreateTargetData().String())
	if optimized {
		passManager := l.NewPassManager()
		passManager.AddAddressSanitizerFunctionPass()
		passManager.AddArgumentPromotionPass()
		passManager.AddConstantMergePass()
		passManager.AddDeadArgEliminationPass()
		passManager.AddAggressiveDCEPass()
		passManager.AddLoopUnrollPass()
		passManager.AddLoopDeletionPass()
		passManager.AddLoopUnswitchPass()
		passManager.AddPromoteMemoryToRegisterPass()
		passManager.AddDemoteMemoryToRegisterPass()
		passManager.AddInstructionCombiningPass()
		passManager.AddScalarReplAggregatesPass()
		passManager.AddFunctionInliningPass()
		passManager.AddGlobalDCEPass()
		passManager.AddCoroElidePass()
		passManager.AddConstantMergePass()
		passManager.AddTailCallEliminationPass()
		passManager.AddStripDeadPrototypesPass()
		passManager.AddCFGSimplificationPass()
		passManager.Run(module)
		targetMachine.AddAnalysisPasses(passManager)
	}
	memBuffer, err := targetMachine.EmitToMemoryBuffer(module, l.ObjectFile)
	if err != nil {
		return "", err
	}
	fdOutput, err := os.Create(path)
	if err != nil {
		return "", err
	}
	_, err = fdOutput.Write(memBuffer.Bytes())
	if err != nil {
		return "", err
	}

	return path, err
}
