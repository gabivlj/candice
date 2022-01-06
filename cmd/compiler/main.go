package main

import (
	"fmt"
	"github.com/gabivlj/candice/internals/compiler"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/internals/semantic"
	"os"
	"strconv"
)

func main() {
	file := os.Args[1]
	code, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error opening file: ", err.Error())
	}

	p := parser.New(lexer.New(string(code)))
	tree := p.Parse()
	if len(p.Errors) > 0 {
		fmt.Println("Error parsing file:")
		for i, err := range p.Errors {
			fmt.Println("Error", strconv.FormatInt(int64(i), 10), err)
		}
		return
	}
	s := semantic.New()
	s.Analyze(tree)
	if len(s.Errors) > 0 {
		fmt.Println("Error analyzing file:")
		for i, err := range s.Errors {
			fmt.Println("Error", strconv.FormatInt(int64(i), 10), err)
		}
		return
	}

	c := compiler.New()
	c.Compile(tree)
	err = c.GenerateExecutable()
	if err != nil {
		fmt.Println("Error compiling code:")
		fmt.Println(err)
	}
}
