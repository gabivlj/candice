package ctypes

import (
	"fmt"
	"strings"
)

/// Candice types

var I8 = &Integer{BitSize: 8}
var I16 = &Integer{BitSize: 16}
var I32 = &Integer{BitSize: 32}
var I64 = &Integer{BitSize: 64}

var typeLiteral = map[string]Type{
	"i8":   I8,
	"i16":  I16,
	"i32":  I32,
	"i64":  I64,
	"void": VoidType,
	"i0":   VoidType,
}

func LiteralToType(literal string) Type {
	if t, ok := typeLiteral[literal]; ok {
		return t
	}
	return &Anonymous{Name: literal}
}

var todoType = &Anonymous{Name: "<TODO>"}

func TODO() Type {
	return todoType
}

// Type is the implementation of a candice type
type Type interface {
	// candiceType private flag
	candiceType()

	// SizeOf returns the size in bytes
	SizeOf() int64

	// Alignment returns the alignment in bytes
	Alignment() int64

	String() string
}

var VoidType = &Void{}

type Void struct{}

func (_ *Void) String() string { return "void" }

func (_ *Void) candiceType() {}

func (_ *Void) SizeOf() int64 {
	return 0
}

func (_ *Void) Alignment() int64 {
	return 0
}

type Integer struct {
	// 8bits, 16bits, 32bits, 64bits
	BitSize uint
}

func (i *Integer) String() string {
	return fmt.Sprintf("i%d", i.BitSize)
}

func (_ *Integer) candiceType() {}

func (i *Integer) SizeOf() int64 {
	return int64(i.BitSize / 8)
}

func (i *Integer) Alignment() int64 {
	return int64(i.BitSize / 8)
}

type UInteger struct {
	// 8bits, 16bits, 32bits, 64bits
	BitSize uint
}

func (i *UInteger) String() string {
	return fmt.Sprintf("u%d", i.BitSize)
}

func (i *UInteger) SizeOf() int64 {
	return int64(i.BitSize / 8)
}

func (i *UInteger) Alignment() int64 {
	return int64(i.BitSize / 8)
}

func (_ *UInteger) candiceType() {}

type Pointer struct {
	Inner Type
}

func (p *Pointer) String() string {
	return "*" + p.Inner.String()
}

func (_ *Pointer) Alignment() int64 {
	return 8
}

func (_ *Pointer) SizeOf() int64 { return 8 }

func (_ *Pointer) candiceType() {}

type Float struct{}

func (_ *Float) candiceType() {}

func (f *Float) SizeOf() int64 {
	return 8
}

func (f *Float) String() string {
	return "f64"
}

func (_ *Float) Alignment() int64 {
	return 8
}

type Array struct {
	Inner  Type
	Length int64
}

func (a *Array) String() string {
	return fmt.Sprintf("[%d]%s", a.Length, a.Inner.String())
}

func (a *Array) SizeOf() int64 {
	return a.Inner.SizeOf() * a.Length
}

func (a *Array) Alignment() int64 {
	return a.Inner.Alignment()
}

func (_ *Array) candiceType() {}

type Function struct {
	Name       string
	Parameters []Type
	Names      []string
	Return     Type
}

func (_ *Function) candiceType() {}

func (f *Function) SizeOf() int64 {
	return 8
}

func (f *Function) Alignment() int64 {
	return 8
}

func (f *Function) FullString() string {
	builder := strings.Builder{}
	builder.WriteString("func ")
	builder.WriteString(f.Name)
	builder.WriteString("(")
	for i := 0; i < len(f.Names); i++ {
		if i >= 1 {
			builder.WriteString(", ")
		}
		builder.WriteString(f.Names[i])
		builder.WriteString(" ")
		builder.WriteString(f.Parameters[i].String())
	}
	builder.WriteString(") ")
	builder.WriteString(f.Return.String())
	return builder.String()
}

func (f *Function) String() string {
	builder := strings.Builder{}
	builder.WriteString("func")
	if f.Name != "" {
		builder.WriteByte(' ')
		builder.WriteString(f.Name)
	}

	builder.WriteString("(")
	for i := 0; i < len(f.Parameters); i++ {
		if i >= 1 {
			builder.WriteString(", ")
		}
		builder.WriteString(f.Parameters[i].String())
	}
	builder.WriteString(")")

	if f.Return != nil {
		builder.WriteByte(' ')
		builder.WriteString(f.Return.String())
	}

	return builder.String()
}

// Anonymous type is a type that is not yet declared or not processed by the semantic tree.
// The front-end compiler will try to lookup by name the type and throw an exception if
// it's not defined. We can do fancy lazy stuff with this.
type Anonymous struct {
	Name    string
	Modules []string
}

func (_ *Anonymous) candiceType()     {}
func (a *Anonymous) String() string   { return strings.Join(append(a.Modules, a.Name), ".") }
func (a *Anonymous) Alignment() int64 { return 0 }
func (a *Anonymous) SizeOf() int64    { return 0 }

type Struct struct {
	Fields []Type
	Names  []string
	Name   string
}

func (s *Struct) GetField(fieldName string) (int, Type) {
	for i, field := range s.Fields {
		if s.Names[i] == fieldName {
			return i, field
		}
	}
	return -1, nil
}

func (s *Struct) FullString() string {
	str := strings.Builder{}
	str.WriteString("struct " + s.Name + " {\n")
	for i, field := range s.Fields {
		if i >= 1 {
			str.WriteByte('\n')
		}
		str.WriteString(fmt.Sprintf("%s %s", s.Names[i], field.String()))
	}
	str.WriteString("\n}")
	return str.String()
}

func (s *Struct) String() string {
	return s.Name
}

func (_ *Struct) candiceType() {}

func (s *Struct) SizeOf() int64 {
	currentAddress := int64(0)
	// The padding formulas are quite confusing at first, check this out!
	// -- https://en.wikipedia.org/wiki/Data_structure_alignment#Computing_padding
	for _, t := range s.Fields {
		typeSize := t.SizeOf()
		alignment := t.Alignment()
		offset := (currentAddress - (currentAddress % alignment)) % alignment
		currentAddress += offset
		currentAddress += (alignment - (currentAddress % alignment)) % alignment
		currentAddress += typeSize
	}
	alignment := s.Alignment()
	currentAddress += (currentAddress - (currentAddress % alignment)) % alignment
	currentAddress += (alignment - (currentAddress % alignment)) % alignment
	return currentAddress
}

func (s *Struct) Alignment() int64 {
	maximumAlignment := int64(0)
	for _, t := range s.Fields {
		alignment := t.Alignment()
		if alignment >= maximumAlignment {
			maximumAlignment = alignment
		}
	}
	return maximumAlignment
}

func IsNumeric(t Type) bool {
	switch t.(type) {
	case *Integer:
		return true
	case *UInteger:
		return true
	case *Float:
		return true
	}
	return false
}

func IsPointer(t Type) bool {
	_, ok := t.(*Pointer)
	return ok
}

func IsArray(t Type) bool {
	_, ok := t.(*Array)
	return ok
}
