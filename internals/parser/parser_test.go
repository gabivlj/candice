package parser

import (
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
