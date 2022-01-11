package program

import (
	"flag"
	"github.com/gabivlj/candice/internals/compiler"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/internals/semantic"
	"github.com/gabivlj/candice/pkg/logger"
	"os"
	"strconv"
	"strings"
	"time"
)

func Init() {
	current := time.Now()
	var objectFileFlag string
	var programName string
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

	c := compiler.New()
	c.Compile(tree)
	err = c.GenerateExecutableExperimental(programName, objects)
	if err != nil {
		logger.Error("Internally At Compile Time", err.Error())
		return
	}
	passedTime := float64(time.Now().UnixMilli() - current.UnixMilli())
	logger.Success("BUILD SUCCESSFUL. (" + strconv.FormatFloat(passedTime/1000, 'f', 3, 64) + "s)")
}
