package tests

import (
	"github.com/gabivlj/candice/internals/compiler"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/internals/semantic"
	"log"
	"os"
	"testing"
)

func TestSrcs(t *testing.T) {
	expectedOutputs := map[string]string{
		"struct.cd":         "-3 0 -3 43",
		"cast.cd":           "32",
		"functions.cd":      "3 4 5 5",
		"ifstmt.cd":         "1 2 3 1 2 3 1 2 3",
		"if_statements2.cd": "4 4 4",
		"if_statements3.cd": "300",
		"for_statement.cd":  "0 1 2 3 4 0 1 2 3 4 0 1 2 3 4",
		"for_statement.cd2": "1 2 3 4 5 6 7 8 9 10",
	}
	elems, err := os.ReadDir("./src")
	if err != nil {
		t.Fatal(err)
	}
	for _, elem := range elems {
		txt, err := os.ReadFile("./src/" + elem.Name())
		if err != nil {
			t.Fatal(err)
		}
		s := semantic.New()
		p := parser.New(lexer.New(string(txt)))
		root := p.Parse()
		if len(p.Errors) > 0 {
			t.Fatal(p.Errors)
		}

		s.Analyze(root)
		if len(s.Errors) > 0 {
			t.Fatal(s.Errors)
		}

		c := compiler.New()
		c.Compile(root)
		output, err := c.Execute()
		expected := expectedOutputs[elem.Name()]
		if expected != string(output) {
			t.Fatal("test for", elem.Name(), "failed, expected output", expected, "got:", string(output), " ", err)
		}
		log.Println("File ", elem.Name(), " passed")
	}
}
