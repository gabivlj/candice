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
			`variable : i64 = @cast(i64, 32 + 32 * 32 * 32 * 44 / 33 / 44 || 444 && 333 ^ 333 && 333 & 44)`,
			true,
		},
		{
			`variable : i32 = !!!*&44
					 variable2 := variable + variable`,
			true,
		},
		{
			`variable : i32 = *!!!*&44`,
			false,
		},
		{
			`struct Point { x i32 y i32 point *Point }`,
			true,
		},
		{
			`struct Point { x i32 y i32 point Point }`,
			false,
		},
		{
			`func aFunction(i i32, j i32) i32 {
	if 1 {
		return 3
	}
	return;
}`,
			false,
		},
		{
			`func aFunction(i i32, j i32) i32 {
	if 1 {
		return i + j;
	} else { return 3; }
	
}`,
			true,
		},
		{
			`func aFunction(i i32, j i32) i32 {
	if 1 {
		return i + j;
	}
}`,
			false,
		},
		{
			`func aFunction(i i32, j i32) {
	if 1 {
		return;
	}
}`,
			true,
		},
		{
			`func aFunction(i i32, j i32) {}`,
			true,
		},
		{
			`for 1 < 100 { }`,
			true,
		},
		{
			`for i := 100; i < 100; { }`,
			true,
		},
		{
			`for i := 100; i < 100; i = i + 1 { }`,
			true,
		},
		{
			`func callMe() i32 { return callMe() }
					called := callMe() + 3 + 10 * 100 / 300;
					func callback(c func() i32) { c(); } callback(callMe)`,
			true,
		},
		{
			`func callMe() i32 { return callMe() }
					called := callMe() + 3 + 10 * 100 / 300;
					func callback(c func() i64) { c(); } callback(callMe)`,
			false,
		},
	}

	for _, test := range tests {
		semantic := New()
		p := parser.New(lexer.New(test.program))
		program := p.Parse()
		a.Assert(len(p.Errors) == 0, p.Errors)
		semantic.Analyze(program)
		if test.shouldBeOk && len(semantic.Errors) != 0 {
			t.Fatal(test, semantic.Errors)
		} else if !test.shouldBeOk && len(semantic.Errors) == 0 {
			t.Fatal(test, "shouldn't be ok but we got 0 Errors...")
		}
	}
}
