package tests

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/gabivlj/candice/internals/compiler"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/internals/semantic"
	"github.com/gabivlj/candice/pkg/logger"
)

func TestSrcs(t *testing.T) {
	expectedOutputs := map[string]string{
		"struct.cd":             "-3 0 -3 43",
		"cast.cd":               "32 -1 -1 -1 -1 0 0 0 0 -10.000 1.000",
		"functions.cd":          "3 4 5 5",
		"ifstmt.cd":             "1 2 3 1 2 3 1 2 3",
		"if_statements2.cd":     "4 4 4",
		"if_statements3.cd":     "300",
		"for_statement.cd":      "0 1 2 3 4 0 1 2 3 4 0 1 2 3 4",
		"for_statement2.cd":     "1 2 3 4 5 6 7 8 9 10",
		"string_literal.cd":     "Hello world!",
		"unsigned_ints.cd":      "4294967200 -96",
		"linked_list.cd":        "100",
		"fibonacci.cd":          "75025",
		"nested_loops.cd":       "0 0 1 2 3 4 5 6 7 8 9 10 1 1 2 3 4 5 6 7 8 9 10 2 2 3 4 5 6 7 8 9 10 3 3 4 5 6 7 8 9 10 4 4 5 6 7 8 9 10 5 5 6 7 8 9 10 6 6 7 8 9 10 7 7 8 9 10 8 8 9 10 9 9 10 10 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100",
		"if_statements4.cd":     "ok nice ok nice",
		"callback.cd":           "10",
		"arrayuser.cd":          "WORKS",
		"blocks.cd":             "11 10 9 8 7 6 5 4 3 2 1",
		"anonymous_function.cd": "100 1",
		"print.cd":              "Thing",
		"macro_if.cd":           fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH),
		"1variadic.cd":          "formatting a number: 3",
		"0arr.cd":               "1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1",
	}

	elems, err := os.ReadDir("./src")

	if err != nil {
		t.Fatal(err)
	}

	for _, elem := range elems {
		if elem.IsDir() {
			continue
		}

		txt, err := os.ReadFile("./src/" + elem.Name())
		if err != nil {
			t.Fatal(err)
		}
		s := semantic.New()
		p := parser.New(lexer.New(string(txt)))
		root := p.Parse()
		logger.WarningNoTag("Running " + elem.Name())
		if len(p.Errors) > 0 {
			t.Fatal(p.Errors)
		}
		semantic.ResetPaths()
		s.ContextDirectoryPath = "./src"
		s.Analyze(root)

		if len(s.Errors) > 0 {
			t.Fatal(s.Errors)
		}

		c := compiler.New(s)
		c.Compile(root)
		output, err := c.Execute()
		if err != nil {
			t.Fatal(err)
		}
		expected := expectedOutputs[elem.Name()]
		if expected != string(output) {
			t.Fatal("test for", elem.Name(), "failed, expected output", expected, "got:", string(output), " ", err)
		}
		logger.Success("File " + elem.Name() + " passed")
	}
}

func BenchmarkSrcs(t *testing.B) {
	expectedOutputs := map[string]string{
		"struct.cd":             "-3 0 -3 43",
		"cast.cd":               "32 -1 -1 -1 -1 0 0 0 0 -10.000 1.000",
		"functions.cd":          "3 4 5 5",
		"ifstmt.cd":             "1 2 3 1 2 3 1 2 3",
		"if_statements2.cd":     "4 4 4",
		"if_statements3.cd":     "300",
		"for_statement.cd":      "0 1 2 3 4 0 1 2 3 4 0 1 2 3 4",
		"for_statement2.cd":     "1 2 3 4 5 6 7 8 9 10",
		"string_literal.cd":     "Hello world!",
		"unsigned_ints.cd":      "4294967200 -96",
		"linked_list.cd":        "100",
		"fibonacci.cd":          "75025",
		"nested_loops.cd":       "0 0 1 2 3 4 5 6 7 8 9 10 1 1 2 3 4 5 6 7 8 9 10 2 2 3 4 5 6 7 8 9 10 3 3 4 5 6 7 8 9 10 4 4 5 6 7 8 9 10 5 5 6 7 8 9 10 6 6 7 8 9 10 7 7 8 9 10 8 8 9 10 9 9 10 10 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100",
		"if_statements4.cd":     "ok nice ok nice",
		"callback.cd":           "10",
		"arrayuser.cd":          "WORKS",
		"blocks.cd":             "11 10 9 8 7 6 5 4 3 2 1",
		"anonymous_function.cd": "100 1",
		"print.cd":              "Thing",
		"macro_if.cd":           fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH),
		"1variadic.cd":          "formatting a number: 3",
		"0arr.cd":               "1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1",
	}

	elems, err := os.ReadDir("./src")

	if err != nil {
		t.Fatal(err)
	}

	for _, elem := range elems {
		if elem.IsDir() {
			continue
		}
		txt, err := os.ReadFile("./src/" + elem.Name())
		if err != nil {
			t.Fatal(err)
		}
		s := semantic.New()
		p := parser.New(lexer.New(string(txt)))
		root := p.Parse()
		logger.WarningNoTag("Running " + elem.Name())
		if len(p.Errors) > 0 {
			t.Fatal(p.Errors)
		}
		semantic.ResetPaths()
		s.ContextDirectoryPath = "./src"
		s.Analyze(root)

		if len(s.Errors) > 0 {
			t.Fatal(s.Errors)
		}

		c := compiler.New(s)
		c.Compile(root)
		output, err := c.Execute()
		if err != nil {
			t.Fatal(err)
		}
		expected := expectedOutputs[elem.Name()]
		if expected != string(output) {
			t.Fatal("test for", elem.Name(), "failed, expected output", expected, "got:", string(output), " ", err)
		}
		logger.Success("File " + elem.Name() + " passed")
	}
}
