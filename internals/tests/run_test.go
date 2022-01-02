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
		"struct.cd": "-3 0 -3",
		"cast.cd":   "32",
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
