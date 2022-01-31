package program

import (
	"flag"
	"fmt"
	"github.com/gabivlj/candice/internals/compiler"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/internals/semantic"
	"github.com/gabivlj/candice/pkg/logger"
	"github.com/gabivlj/candice/pkg/terminal"
	"os"
	"strconv"
	"strings"
	"time"
)

func Init() {
	if !terminal.ClangExists() {
		logger.Warning("It seems like clang doesn't exist in your machine, consider installing it!")
	}

	current := time.Now()
	var objectFileFlag string
	var programName string
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		_, _ = fmt.Fprintf(os.Stderr, "Put the entry point of your candice file at the end of the command\n")
		flag.PrintDefaults()
	}
	flag.StringVar(&objectFileFlag, "o", "", "Objects to link")
	flag.StringVar(&programName, "name", "program", "Program name")
	flag.Parse()
	file := flag.Args()[0]
	objects := strings.Split(objectFileFlag, ",")
	code, err := os.ReadFile(file)
	if err != nil {
		logger.Error("", "Opening file", err.Error())
	}

	p := parser.New(lexer.New(string(code)))
	tree := p.Parse()
	if len(p.Errors) > 0 {
		for _, err := range p.Errors {
			logger.Error("Parsing", err.Error())
		}
		return
	}
	s := semantic.New()
	s.Analyze(tree)
	if len(s.Errors) > 0 {
		for _, err := range s.Errors {
			logger.Error("Analyzing", err.Error())
		}
		return
	}

	c := compiler.New(s)
	c.Compile(tree)
	err = c.GenerateExecutableExperimental(programName, objects)
	if err != nil {
		logger.Error("Internally At Compile Time", err.Error())
		return
	}
	passedTime := float64(time.Now().UnixMilli() - current.UnixMilli())
	logger.Success("BUILD SUCCESSFUL. (" + strconv.FormatFloat(passedTime/1000, 'f', 3, 64) + "s)")
}
