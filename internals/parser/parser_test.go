package parser

import (
	"fmt"
	"log"
	"testing"

	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/pkg/a"
	"github.com/gabivlj/candice/pkg/logger"
)

func TestParser_ParseType(t *testing.T) {
	src := "****[1000][1000][1000]*****[10][10][1302][12221]*************************[1222]i32"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.parseType()
	a.Assert(len(p.Errors) == 0, p.Errors)
	a.AssertEqual(src, tt.String())
}

func TestParser_ParseTypeI32(t *testing.T) {
	src := "i32"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.parseType()
	a.Assert(len(p.Errors) == 0, p.Errors)
	a.AssertEqual(src, tt.String())
	a.Assert(tt.(*ctypes.Integer).BitSize == 32)
}

func TestParser_ParseTypeI64(t *testing.T) {
	src := "i64"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.parseType()
	a.Assert(len(p.Errors) == 0, p.Errors)
	a.AssertEqual(src, tt.String())
	a.Assert(tt.(*ctypes.Integer).BitSize == 64)
}

func TestParser_ParseTypeI16(t *testing.T) {
	src := "i16"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.parseType()
	a.Assert(len(p.Errors) == 0, p.Errors)
	a.AssertEqual(src, tt.String())
	a.Assert(tt.(*ctypes.Integer).BitSize == 16)
}

func TestParser_ParseTypeI8(t *testing.T) {
	src := "i8"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.parseType()
	a.Assert(len(p.Errors) == 0, p.Errors)
	a.AssertEqual(src, tt.String())
	a.Assert(tt.(*ctypes.Integer).BitSize == 8)
}

func TestParser_ParseBinaryOperation(t *testing.T) {
	src := "3+3"
	lex := lexer.New(src)
	p := New(lex)
	tt := p.Parse()
	a.AssertEqual(tt.String(), "(3 + 3);\n")
}

func TestParser_MultipleExpressions(t *testing.T) {
	tests := []struct {
		expression string
		expected   string
	}{
		{
			expression: `if aConstant2 as i64 != aConstant as i64 @print("bad...")`,
			expected: `if (aConstant2 as i64 != aConstant as i64) {
@print("bad...");
}
`,
		},
		{
			expression: `union Union {
					element i32
					otherElement u32
				}`,
			expected: `union Union {
element i32
otherElement u32
}
`,
		},
		{
			expression: "exp :int = 3==3+(3+3)/4*6+-&element.element.element.element",
			expected:   "exp :int = (3 == ((3 + (((3 + 3) / 4) * 6)) + -&element.element.element.element));\n",
		},
		{
			expression: "hello :int = 4; 55  >=   33 && 33 <= 11 || 44==-33",
			expected:   "hello :int = 4;\n(((55 >= 33) && (33 <= 11)) || (44 == -33));\n",
		},
		{
			expression: "*hello.ss[3].ss = 3;",
			expected:   "*hello.ss[3].ss = 3;\n",
		},
		{
			expression: "cool = \"Hello world\"",
			expected:   "cool = \"Hello world\";\n",
		},
		{
			expression: "@print(\"hello world\");",
			expected:   "@print(\"hello world\");\n",
		},
		{
			expression: "pointer : *i32 = @alloc(i32, 33);",
			expected:   "pointer :*i32 = @alloc(i32, 33);\n",
		},
		{
			expression: "pointer : *i32 = @alloc(i32, @alloc(i64, 223 * 333 + 212921) + 329323 & 3333);",
			expected:   "pointer :*i32 = @alloc(i32, (@alloc(i64, ((223 * 333) + 212921)) + (329323 & 3333)));\n",
		},
		{
			expression: `
						if 3 {
						} else if 3+3+3==3/4/3/3.0+hello.hello == 4 {
						} else {
						}
					`,
			expected: `if 3 {

} else if ((((3 + 3) + 3) == ((((3 / 4) / 3) / 3) + hello.hello)) == 4) {

} else {

}
`,
		},
		{
			expression: `
						if 1 @print("hello world");
						else if 2 @print("hello world 2");
						else if 3 @print("hello world 3");
						else @print("hello world else"); @print("hello world apart");
					`,
			expected: `if 1 {
@print("hello world");
} else if 2 {
@print("hello world 2");
} else if 3 {
@print("hello world 3");
} else {
@print("hello world else");
}
@print("hello world apart");
`,
		},
		{
			expression: `if 1 == 1 @print("cool") else @print("not cool")`,
			expected: `if (1 == 1) {
@print("cool");
} else {
@print("not cool");
}
`,
		},
		{
			expression: `for i := 0; i < 1000; i = i + 1 { @print("hello world!") }`,
			expected: `for i : = 0; (i < 1000); i = (i + 1); {
@print("hello world!");
}
`,
		},
		{
			expression: `for i := 0; i < 1000; i = i + 1 @print("hello world!") @print("hello world...")`,
			expected: `for i : = 0; (i < 1000); i = (i + 1); {
@print("hello world!");
}
@print("hello world...");
`,
		},
		{
			expression: `for i.i.i.i.i[0] = 0; i < 1000 && cool && thing || works == 3; // sdds
									i.i.i.i.i[0] = i + 1 @print("hello world!");///sss
// sss`,
			expected: `for i.i.i.i.i[0] = 0; ((((i < 1000) && cool) && thing) || (works == 3)); i.i.i.i.i[0] = (i + 1); {
@print("hello world!");
}
`,
		},
		{
			expression: `for @print("infinite loop")`,
			expected: `for {
@print("infinite loop");
}
`,
		},
		{
			expression: `for { @print("infinite loop") }`,
			expected: `for {
@print("infinite loop");
}
`,
		},
		{
			// NOTE! for 1 @print("infinite loop") cannot be represented like this
			expression: `for 1 { @print("infinite loop") }`,
			expected: `for 1 {
@print("infinite loop");
}
`,
		},
		{
			// NOTE! for 1 @print("infinite loop") is represented like this
			expression: `for 1 @print("infinite loop")`,
			expected: `for {
1;
}
@print("infinite loop");
`,
		},
		{
			expression: `for i := 0; i < 100; {}`,
			expected: `for i : = 0; (i < 100); {

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
			expression: `structLiteral := @StructLiteral{ a: 1, b: &*&@AnotherStruct { pog: 3, pog2: 4, } }`,
			expected: `structLiteral : = @StructLiteral{
a: 1,
b: &*&@AnotherStruct{
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
		{
			expression: `func hello_world(hello i32, world [111]i32) i32 {
						@print(hello + world[0])
						return 0;
					}`,
			expected: `func hello_world(hello i32, world [111]i32) i32 {
@print((hello + world[0]));
return 0;
}
`,
		},
		{
			expression: `func helloworld(hello i32, world [111]a_module.Element) i32 {
				if 33 > hello {
					return module.@hello{hellobaby: 3}
				}
				helloworld(hello + world[0], 0)
				return 0;
			}`,
			expected: `func helloworld(hello i32, world [111]a_module.Element) i32 {
if (33 > hello) {
return module.@hello{
hellobaby: 3,
};
}
helloworld((hello + world[0]), 0);
return 0;
}
`,
		},
		{
			expression: `import hellomodules, i32, i32, i32, "./path.cd";
	break`,
			expected: `import hellomodules, i32, i32, i32, "./path.cd"
break
`,
		},
		{
			expression: `arr := [1000]i32{1, 1, 1, 1, 2}`,
			expected: `arr : = [1000]i32 {1, 1, 1, 1, 2};
`,
		},
		{
			expression: `extern func thing(i32, i32, i32);`,
			expected: `extern func thing(i32, i32, i32) void;
`,
		},
		{
			expression: `extern func thing(i32, i32, i32) i32;`,
			expected: `extern func thing(i32, i32, i32) i32;
`,
		},
		{
			expression: `switch 3 {
				case 3 {
					@print("it's three");
				}

				case 4 {
					@print("it's four");
					@print("it's four!!!!");
				}

				default {
					@print("whatever");
				}
			}`,
			expected: `switch 3 {
case 3 {
@print("it's three");
}
case 4 {
@print("it's four");
@print("it's four!!!!");
}
default {
@print("whatever");
}
}
`,
		},
		{
			expression: `express, express2 := thing(), ole;`,
			expected: `express, express2 : = thing(), ole;
`,
		},
		{
			expression: `ole, ola, olop = 3, 3, 3`,
			expected: `ole, ola, olop = 3, 3, 3;
`,
		},
	}
	for _, test := range tests {
		logger.Warning("running " + test.expression)
		evaluate(t, test.expression, test.expected)
	}
}

func evaluate(t *testing.T, expression, expected string) {
	p := New(lexer.New(expression))
	program := p.Parse()
	if len(p.Errors) != 0 {
		log.Println("error on expression ", expression)
		panic(fmt.Sprintf("%v", p.Errors))
	}
	output := program.String()
	a.AssertEqual(output, expected)
}
