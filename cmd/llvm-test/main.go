// This example produces LLVM IR code equivalent to the following C code, which
// implements a pseudo-random number generator.
//
//    int abs(int x);
//
//    int seed = 0;
//
//    // ref: https://en.wikipedia.org/wiki/Linear_congruential_generator
//    //    a = 0x15A4E35
//    //    c = 1
//    int rand(void) {
//       seed = seed*0x15A4E35 + 1;
//       return abs(seed);
//    }
package main

import (
	"fmt"
	"github.com/gabivlj/candice/internals/compiler"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"log"
)

func main() {
	// Create convenience types and constants.
	i32 := types.I32
	zero := constant.NewInt(i32, 0)
	one := constant.NewInt(i32, 1)
	a := constant.NewInt(i32, 0x15A4E35) // multiplier of the PRNG.
	c := constant.NewInt(i32, 1)         // increment of the PRNG.

	// Create a new LLVM IR module.
	m := ir.NewModule()

	printf := m.NewFunc(
		"printf",
		types.I32,
		ir.NewParam("", types.NewPointer(types.I8)),
	)

	printf.Sig.Variadic = true

	printCool := m.NewFunc(
		"printCool",
		types.Void,
	)

	// Create an external function declaration and append it to the module.
	//
	//    int abs(int x);
	abs := m.NewFunc("abs", i32, ir.NewParam("x", i32))

	// Create a global variable definition and append it to the module.
	//
	//    int seed = 0;
	seed := m.NewGlobalDef("seed", constant.NewInt(i32, 12))
	log.Println(seed.Type())
	// Create a function definition and append it to the module.
	//
	//    int rand(void) { ... }
	rand := m.NewFunc("rand", i32)

	// Create an unnamed entry basic block and append it to the `rand` function.
	entry := rand.NewBlock("")

	// Create instructions and append them to the entry basic block.
	tmp1 := entry.NewLoad(i32, seed)
	tmp2 := entry.NewMul(tmp1, a)
	tmp3 := entry.NewAdd(tmp2, c)
	entry.NewStore(tmp3, seed)
	tmp4 := entry.NewCall(abs, tmp3)
	entry.NewRet(tmp4)
	fun := m.NewFunc("main", i32)
	entryBlock := fun.NewBlock("")
	value := entryBlock.NewCall(rand)
	s := "Result: %d\n"
	printInteger := m.NewGlobalDef("printinteger", constant.NewCharArrayFromString(s))
	log.Println(printInteger.Type())
	pointerToString := entryBlock.NewGetElementPtr(types.NewArray(uint64(len(s)), types.I8), printInteger, zero, zero)
	log.Println(pointerToString.Type())
	malloc := m.NewFunc("malloc", types.NewPointer(types.I8), ir.NewParam("size", types.I64))
	pointer := entryBlock.NewCall(malloc, constant.NewInt(types.I64, 4 * 4))
	i32Pointer := entryBlock.NewBitCast(pointer, types.NewPointer(types.I32))
	allocaPointer := entryBlock.NewAlloca(i32Pointer.Type())
	entryBlock.NewStore(i32Pointer, allocaPointer)
	log.Println(allocaPointer.Type())
	addr := entryBlock.NewLoad(types.NewPointer(types.I32), allocaPointer)
	entryBlock.NewStore(constant.NewInt(types.I32, 32), addr)
	element :=
		entryBlock.NewGetElementPtr(
			addr.Type().(*types.PointerType).ElemType, addr, zero,
		)
	log.Println(element)
	entryBlock.NewCall(printf, pointerToString, entryBlock.NewLoad(types.I32, element))
	//entryBlock.NewLoad(types.I32, pointerFirst)
	entryBlock.NewCall(printf, pointerToString, value)
	entryBlock.NewCall(printCool)

	struc := types.NewStruct(types.I32, types.I32)
	instance := entryBlock.NewAlloca(struc)
	instanceFirst := entryBlock.NewGetElementPtr(struc, instance, zero, zero)
	entryBlock.NewStore(constant.NewInt(types.I32, 43), instanceFirst)
	instanceSecond := entryBlock.NewGetElementPtr(struc, instance, zero, one)
	entryBlock.NewStore(constant.NewInt(types.I32, 3223), instanceSecond)
	coolInteger := entryBlock.NewLoad(types.I32, instanceFirst)
	entryBlock.NewCall(printf, pointerToString, coolInteger)
	coolInteger2 := entryBlock.NewLoad(types.I32, instanceSecond)
	entryBlock.NewCall(printf, pointerToString, coolInteger2)
	entryBlock.NewRet(constant.NewInt(i32, 0))
	fmt.Println(compiler.GenerateExecutable(m, "exec"))
}