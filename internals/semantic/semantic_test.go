package semantic

import (
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/pkg/a"
	"testing"
)

func TestSemantic_Analyze(t *testing.T) {
	tests := []struct {
		program    string
		shouldBeOk bool
	}{
		{
			`variable : i64 = @cast(i64, 32); variable : i32 = 32`,
			true,
		},
		{
			`variable : i32 = @cast(i64, 32);`,
			false,
		},
		{
			`variable : i64 = 32`,
			false,
		},
		{
			`variable : i64 = @cast(i64, 32 + 32 * 32 * 32 * 44)`,
			true,
		},
	}

	for _, test := range tests {
		semantic := New()
		p := parser.New(lexer.New(test.program))
		program := p.Parse()
		a.Assert(len(p.Errors) == 0)
		semantic.Analyze(program)
		if test.shouldBeOk && len(semantic.errors) != 0 {
			t.Fatal(test, semantic.errors)
		} else if !test.shouldBeOk && len(semantic.errors) == 0 {
			t.Fatal(test, "shouldn't be ok but we got 0 errors...")
		}
	}

}
