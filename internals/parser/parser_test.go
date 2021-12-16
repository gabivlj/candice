package parser

import (
	"fmt"
	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/pkg/a"
	"testing"
)

func TestParser_ParseType(t *testing.T) {
	src := "****[1000][1000][1000]*****[10][10][1302][12221]*************************[1222]i32"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.parseType()
	a.Assert(len(p.errors) == 0, p.errors)
	a.AssertEqual(src, tt.String())
}

func TestParser_ParseTypeI32(t *testing.T) {
	src := "i32"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.parseType()
	a.Assert(len(p.errors) == 0, p.errors)
	a.AssertEqual(src, tt.String())
	a.Assert(tt.(*ctypes.Integer).BitSize == 32)
}

func TestParser_ParseTypeI64(t *testing.T) {
	src := "i64"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.parseType()
	a.Assert(len(p.errors) == 0, p.errors)
	a.AssertEqual(src, tt.String())
	a.Assert(tt.(*ctypes.Integer).BitSize == 64)
}

func TestParser_ParseTypeI16(t *testing.T) {
	src := "i16"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.parseType()
	a.Assert(len(p.errors) == 0, p.errors)
	a.AssertEqual(src, tt.String())
	a.Assert(tt.(*ctypes.Integer).BitSize == 16)
}

func TestParser_ParseTypeI8(t *testing.T) {
	src := "i8"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.parseType()
	a.Assert(len(p.errors) == 0, p.errors)
	a.AssertEqual(src, tt.String())
	a.Assert(tt.(*ctypes.Integer).BitSize == 8)
}

func TestParser_ParseBinaryOperation(t *testing.T) {
	src := "3+3"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.Parse()
	a.AssertEqual(tt.String(), "(3+3);\n")
}

func TestParser_MultipleExpressions(t *testing.T) {
	tests := []struct {
		expression string
		expected   string
	}{
		{
			expression: "exp :int = 3==3+(3+3)/4*6+-&element.element.element.element",
			expected:   "exp :int = (3==((3+(((3+3)/4)*6))+-&(((element.element).element).element)));\n",
		},
		{
			expression: "hello :int = 4; 55  >=   33 && 33 <= 11 || 44==-33",
			expected:   "hello :int = 4;\n(((55>=33)&&(33<=11))||(44==-33));\n",
		},
		{
			expression: "*hello.ss[3].ss = 3;",
			expected:   "*((hello.ss[3]).ss) = 3;\n",
		},
		{
			expression: "cool = \"Hello world\"",
			expected:   "cool = \"Hello world\";\n",
		},
		{
			expression: "@println(\"hello world\");",
			expected:   "@println(\"hello world\");\n",
		},
		{
			expression: "pointer : *i32 = @alloc(i32, 33);",
			expected:   "pointer :*i32 = @alloc(i32, 33);\n",
		},
		{
			expression: "pointer : *i32 = @alloc(i32, @alloc(i64, 223 * 333 + 212921) + 329323 & 3333);",
			expected:   "pointer :*i32 = @alloc(i32, (@alloc(i64, ((223*333)+212921))+(329323&3333)));\n",
		},
		{
			expression: `
				if 3 {
				} else if 3+3+3==3/4/3/3.hello.hello == 4 {
				} else {
				}
			`,
			expected: `if 3 {

} else if ((((3+3)+3)==(((3/4)/3)/((3.hello).hello)))==4) {

} else {

}
`,
		},
		{
			expression: `
				if 1 @println("hello world"); 
				else if 2 @println("hello world 2");
				else if 3 @println("hello world 3");
				else @println("hello world else"); @println("hello world apart"); 
			`,
			expected: `if 1 {
@println("hello world");
} else if 2 {
@println("hello world 2");
} else if 3 {
@println("hello world 3");
} else {
@println("hello world else");
}
@println("hello world apart");
`,
		},
		{
			expression: `if 1 == 1 @println("cool") else @println("not cool")`,
			expected: `if (1==1) {
@println("cool");
} else {
@println("not cool");
}
`,
		},
		{
			expression: `for i := 0; i < 1000; i = i + 1 { @println("hello world!") }`,
			expected: `for i :<TODO> = 0; (i<1000); i = (i+1); {
@println("hello world!");
}
`,
		},
		{
			expression: `for i := 0; i < 1000; i = i + 1 @println("hello world!") @println("hello world...")`,
			expected: `for i :<TODO> = 0; (i<1000); i = (i+1); {
@println("hello world!");
}
@println("hello world...");
`,
		},
		{
			expression: `for i.i.i.i.i[0] = 0; i < 1000 && cool && thing || works == 3;
							i.i.i.i.i[0] = i + 1 @println("hello world!");`,
			expected: `for ((((i.i).i).i).i[0]) = 0; ((((i<1000)&&cool)&&thing)||(works==3)); ((((i.i).i).i).i[0]) = (i+1); {
@println("hello world!");
}
`,
		},
		{
			expression: `for @println("infinite loop")`,
			expected: `for {
@println("infinite loop");
}
`,
		},
		{
			expression: `for { @println("infinite loop") }`,
			expected: `for {
@println("infinite loop");
}
`,
		},
		{
			// NOTE! for 1 @println("infinite loop") cannot be represented like this
			expression: `for 1 { @println("infinite loop") }`,
			expected: `for 1 {
@println("infinite loop");
}
`,
		},
		{
			// NOTE! for 1 @println("infinite loop") is represented like this
			expression: `for 1 @println("infinite loop")`,
			expected: `for {
1;
}
@println("infinite loop");
`,
		},
		{
			expression: `for i := 0; i < 100; {}`,
			expected: `for i :<TODO> = 0; (i<100); {

}
`,
		},
		{
			expression: `
				struct Point {
					x i32
					y i32
					p ********[100][100][100]OtherStruct
					
				}
			`,
			expected: `struct Point {
x i32
y i32
p ********[100][100][100]OtherStruct
}
`,
		},
		{
			expression: `structLiteral := StructLiteral{ a: 1, b: &*&AnotherStruct { pog: 3, pog2: 4, } }`,
			expected: `structLiteral :<TODO> = StructLiteral{
a: 1,
b: &*&AnotherStruct{
pog: 3,
pog2: 4,
},
};
`,
		},
		{
			expression: `structLiteral : func (i32, i32) i32 = function`,
			expected: `structLiteral :func(i32, i32) i32 = function;
`,
		},
	}
	for _, test := range tests {
		evaluate(t, test.expression, test.expected)
	}
}

func evaluate(t *testing.T, expression, expected string) {
	p := New(lexer.New(expression))
	program := p.Parse()
	if len(p.errors) != 0 {
		panic(fmt.Sprintf("%v", p.errors))
	}
	output := program.String()
	a.AssertEqual(output, expected)
}
