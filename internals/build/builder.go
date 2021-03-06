package build

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"

	paths "path"

	"github.com/gabivlj/candice/internals/compiler"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/internals/semantic"
	"github.com/gabivlj/candice/internals/tree_printer"
	"github.com/gabivlj/candice/pkg/logger"
)

type Project struct {
}

func ExecuteProject() {
	defer func() {
		// output := recover()
		// if output != nil {
		// 	logger.Error("Internal Compiler Panic", "There has been an internal panic on the compiler", "\n", output)
		// }
	}()

	current := time.Now()
	flags, err := retrieveFlags()
	if err != nil {
		logger.Error("Flags", "Error retrieving flags: "+err.Error(), `
	Usage:
		candice <mode> <path> <flags>
	Modes:
		run - Run the project in the desired path.
		build - Creates an executable of the project in the desired path.
		init - Creates a candice project
		tree - Showcases an AST of the file in the terminal
	Flags:
		--release - Create or runs an optimized build of the project (run, build).
		`)
		return
	}

	if flags.Mode == "init" {
		createSampleProject(flags.Path)
		return
	}

	if flags.Mode == "tree" {
		bytes, err := os.ReadFile(flags.Path)
		if err != nil {
			logger.Error("File", err.Error())
			return
		}

		if err := tree_printer.WriteOutput(string(bytes), os.Stdout); err != nil {
			logger.Error("Generating AST", err.Error())
		}
		return
	}

	config, err := ParseConfigurationFile(paths.Join(flags.Path, "candice.json"))

	if err != nil {
		logger.Error("Project", err.Error())
		return
	}

	codeEntryPoint, err := os.ReadFile(config.EntryPoint)

	if err != nil {
		logger.Error("Project", err.Error())
		return
	}

	p := parser.New(lexer.New(string(codeEntryPoint)))
	tree := p.Parse()
	if len(p.Errors) > 0 {
		for _, err := range p.Errors {
			logger.Error("Parsing", err.Error())
		}
		return
	}
	s := semantic.New()
	s.Analyze(tree)

	for _, err := range s.Warnings {
		logger.Warning(err.Error())
	}

	if len(s.Errors) > 0 {
		for _, err := range s.Errors {
			logger.Error("Analyzing", err.Error())
		}
		return
	}

	c := compiler.New(s)
	c.CompileWithEventHandler(tree, func(e compiler.Event) {
		if e.Kind == compiler.AddFlags {
			config.CompilerFlags = append(config.CompilerFlags, e.Data)
		}
	})

	if flags.Release {
		config.CompilerFlags = append(config.CompilerFlags, "-O3")
	}

	if config.BinaryKind == Object {
		config.CompilerFlags = append(config.CompilerFlags, "-c")
		config.Output += ".o"
	}

	if config.CompileKind == PureLLVM {
		err := c.GenerateExecutableExperimental(config.Output, config.CXX, config.CompilerFlags, flags.Release, config.BinaryKind != Object)
		if err != nil {
			logger.Error("Internally At Compile Time", err.Error())
			return
		}

	} else if config.CompileKind == CXX {
		err := c.GenerateExecutableCXX(config.Output, config.CXX, config.CompilerFlags)
		if err != nil {
			logger.Error("Internally At Compile Time", err.Error())
			return
		}

	} else {
		logger.Error("Configuration", "Unknown compiling kind, use either 'llvm' or 'cxx'.")
		return
	}

	if flags.Mode == "run" {
		if config.BinaryKind == Object {
			logger.Error("you can't run a project that needs to an object!", "Consider setting 'binary' to 'exe' in candice.json")
			return
		}

		cmd := exec.Command("./" + config.Output)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		passedTime := float64(time.Now().UnixMilli() - current.UnixMilli())
		logger.Success("BUILD SUCCESSFUL. (" + strconv.FormatFloat(passedTime/1000, 'f', 3, 64) + "s)")
	}
}

func createSampleProject(basePath string) {

	relative, err := os.Getwd()
	if err != nil {
		logger.Warning("internal error retrieving relative path: " + err.Error())
	}

	candiceJson := path.Join(relative, basePath, "candice.json")
	fd, err := os.Create(candiceJson)
	if err != nil {
		logger.Error("Setup", "error setting up candice.json", err.Error())
		return
	}

	_, err = fd.Write([]byte(CandiceJSONDefault))
	if err != nil {
		logger.Error("Setup", "error setting up candice.json", err.Error())
		return
	}

	candiceMain := path.Join(relative, basePath, "main.cd")
	fdMain, err := os.Create(candiceMain)
	if err != nil {
		logger.Error("Setup", "error setting up main.cd", err.Error())
		return
	}

	_, err = fdMain.Write([]byte(CandiceProgramDefault))
	if err != nil {
		logger.Error("Setup", "error setting up main.cd", err.Error())
		return
	}

	logger.Success("Project created on path " + candiceJson + " " + "\nEnjoy coding!")
}
