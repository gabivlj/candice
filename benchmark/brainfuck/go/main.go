/*
 Brainfuck-Go ( http://github.com/kgabis/brainfuck-go )
 Copyright (c) 2013 Krzysztof Gabis
 Permission is hereby granted, free of charge, to any person obtaining a copy
 of this software and associated documentation files (the "Software"), to deal
 in the Software without restriction, including without limitation the rights
 to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 copies of the Software, and to permit persons to whom the Software is
 furnished to do so, subject to the following conditions:
 The above copyright notice and this permission notice shall be included in
 all copies or substantial portions of the Software.
 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 THE SOFTWARE.
*/
package main

//./brainfuck  2.10s user 0.09s system 99% cpu 2.194 total
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Instruction struct {
	operator uint16
	operand  uint16
}

const (
	op_inc_dp = iota
	op_dec_dp
	op_inc_val
	op_dec_val
	op_out
	op_in
	op_jmp_fwd
	op_jmp_bck
)

const data_size int = 65535

func compile_bf(input io.Reader) (program []Instruction, err error) {
	var pc, jmp_pc uint16 = 0, 0
	jmp_stack := make([]uint16, 0)
	c := make([]byte, 1)
	for {
		n, err := input.Read(c[:])
		if err != nil || n == 0 {
			break
		}
		switch c[0] {
		case '>':
			program = append(program, Instruction{op_inc_dp, 0})
		case '<':
			program = append(program, Instruction{op_dec_dp, 0})
		case '+':
			program = append(program, Instruction{op_inc_val, 0})
		case '-':
			program = append(program, Instruction{op_dec_val, 0})
		case '.':
			program = append(program, Instruction{op_out, 0})
		case ',':
			program = append(program, Instruction{op_in, 0})
		case '[':
			program = append(program, Instruction{op_jmp_fwd, 0})
			jmp_stack = append(jmp_stack, pc)
		case ']':
			if len(jmp_stack) == 0 {
				return nil, errors.New("Compilation error.")
			}
			jmp_pc = jmp_stack[len(jmp_stack)-1]
			jmp_stack = jmp_stack[:len(jmp_stack)-1]
			program = append(program, Instruction{op_jmp_bck, jmp_pc})
			program[jmp_pc].operand = pc
		default:
			pc--
		}
		pc++
	}
	if len(jmp_stack) != 0 {
		return nil, errors.New("Compilation error.")
	}
	return
}

func execute_bf(program []Instruction) error {
	data := make([]int16, data_size)
	var data_ptr uint16 = 0
	reader := bufio.NewReader(os.Stdin)
	for pc := 0; pc < len(program); pc++ {
		switch program[pc].operator {
		case op_inc_dp:
			data_ptr++
		case op_dec_dp:
			data_ptr--
		case op_inc_val:
			data[data_ptr]++
		case op_dec_val:
			data[data_ptr]--
		case op_out:
			// fmt.Printf("%c", data[data_ptr])
			//
		case op_in:
			read_val, _ := reader.ReadByte()
			data[data_ptr] = int16(read_val)
		case op_jmp_fwd:
			if data[data_ptr] == 0 {
				pc = int(program[pc].operand)
			}
		case op_jmp_bck:
			if data[data_ptr] > 0 {
				pc = int(program[pc].operand)
			}
		default:
			return errors.New("error.")
		}
	}
	return nil
}

func main() {
	filename := "../program.b"
	fileContents, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error reading %s\n", filename)
		return
	}
	program, err := compile_bf(fileContents)
	fileContents.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10000; i++ {
		if err := execute_bf(program); err != nil {
			panic(err)
		}
	}
}
