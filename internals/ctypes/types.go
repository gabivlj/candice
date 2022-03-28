package ctypes

import (
	"fmt"
	"strings"

	"github.com/gabivlj/candice/internals/helper"
)

/// Candice types

var I64 = &Integer{BitSize: 64}
var I32 = &Integer{BitSize: 32}
var I16 = &Integer{BitSize: 16}
var I8 = &Integer{BitSize: 8}
var I1 = &Integer{BitSize: 1}
var U64 = &UInteger{BitSize: 64}
var U32 = &UInteger{BitSize: 32}
var U16 = &UInteger{BitSize: 16}
var U8 = &UInteger{BitSize: 8}
var F64 = &Float{BitSize: 64}
var F32 = &Float{BitSize: 32}

// NOTE: I don't think this should be possible
//var F16 = &Float{BitSize: 16}
//var F8 = &Float{BitSize: 8}

var typeLiteral = map[string]Type{
	"i1":   I1,
	"i8":   I8,
	"i16":  I16,
	"i32":  I32,
	"i64":  I64,
	"u8":   U8,
	"u16":  U16,
	"u32":  U32,
	"u64":  U64,
	"f32":  F32,
	"f64":  F64,
	"void": VoidType,
	"i0":   VoidType,
	"bool": I1,
}

func LiteralToType(literal string) Type {
	if t, ok := typeLiteral[literal]; ok {
		return t
	}
	return nil
}

var todoType = &Anonymous{Name: "errors"}

func TODO() Type {
	return todoType
}

// Type is the implementation of a candice type
type Type interface {
	CandiceType()

	// SizeOf returns the size in bytes
	SizeOf() int64

	// Alignment returns the alignment in bytes
	Alignment() int64

	String() string
}

var VoidType = &Void{}

type Void struct{}

func (_ *Void) String() string { return "void" }

func (_ *Void) CandiceType() {}

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

func (_ *Integer) CandiceType() {}

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

func (_ *UInteger) CandiceType() {}

type Pointer struct {
	Inner Type
}

func (p *Pointer) UnwrapAndDepth() (Type, int) {
	var t Type = p
	i := 0
	for IsPointer(t) {
		t = t.(*Pointer).Inner
		i++
	}
	return t, i
}

func (p *Pointer) String() string {
	return "*" + p.Inner.String()
}

func (_ *Pointer) Alignment() int64 {
	return 8
}

func (_ *Pointer) SizeOf() int64 { return 8 }

func (_ *Pointer) CandiceType() {}

type Float struct {
	BitSize uint
}

func (_ *Float) CandiceType() {}

func (f *Float) SizeOf() int64 {
	return int64(f.BitSize / 8)
}

func (f *Float) String() string {
	return fmt.Sprintf("f%d", f.BitSize)
}

func (f *Float) Alignment() int64 {
	return int64(f.BitSize / 8)
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

func (_ *Array) CandiceType() {}

type Function struct {
	Name                     string
	ExternalName             string
	RedefineWithOriginalName bool
	InfiniteParameters       bool
	Names                    []string
	Parameters               []Type
	Return                   Type
}

func (f *Function) IsMainFunction() bool {
	return helper.RetrieveID(f.Name) == "main"
}

func (_ *Function) CandiceType() {}

func (f *Function) SizeOf() int64 {
	return 8
}

func (f *Function) Alignment() int64 {
	return 8
}

func (f *Function) FullString() string {
	builder := strings.Builder{}
	builder.WriteString("func ")
	if f.Name != "" {
		builder.WriteString(helper.RetrieveID(f.Name))
	}
	builder.WriteString("(")
	for i := 0; i < len(f.Names); i++ {
		if i >= 1 {
			builder.WriteString(", ")
		}
		builder.WriteString(helper.RetrieveID(f.Names[i]))
		builder.WriteString(" ")
		builder.WriteString(f.Parameters[i].String())
	}
	builder.WriteString(") ")
	builder.WriteString(f.Return.String())
	return builder.String()
}

func (f *Function) String() string {
	builder := strings.Builder{}
	visualName := helper.RetrieveID(f.Name)
	builder.WriteString("func")
	if visualName != "" {
		builder.WriteByte(' ')
		builder.WriteString(visualName)
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

type FieldType interface {
	Type
	GetName() string
	candiceType()
	GetField(string) (int, Type)
	FullString() string
}

type Union struct {
	Fields []Type
	Names  []string
	Name   string
	ID     string
}

func (u *Union) candiceType() {}

func (u *Union) GetName() string {
	return u.Name
}

func (u *Union) GetField(fieldName string) (int, Type) {
	for i, field := range u.Fields {
		if u.Names[i] == fieldName {
			return i, field
		}
	}
	return -1, nil
}

func (u *Union) FullString() string {
	str := strings.Builder{}
	str.WriteString("union " + strings.Split(u.Name, "-")[0] + " {\n")
	for i, field := range u.Fields {
		if i >= 1 {
			str.WriteByte('\n')
		}
		str.WriteString(fmt.Sprintf("%s %s", u.Names[i], field.String()))
	}
	str.WriteString("\n}")
	return str.String()
}

func (u *Union) String() string {
	return strings.Split(u.Name, "-")[0]
}

func (_ *Union) CandiceType() {}

func (u *Union) SizeOf() int64 {
	max := int64(0)
	for _, t := range u.Fields {
		typeSize := t.SizeOf()
		if max < typeSize {
			max = typeSize
		}
	}

	return max
}

func (u *Union) Alignment() int64 {
	maximumAlignment := int64(0)
	for _, t := range u.Fields {
		alignment := t.Alignment()
		if alignment > maximumAlignment {
			maximumAlignment = alignment
		}
	}

	return maximumAlignment
}

// Anonymous type is a type that is not yet declared or not processed by the semantic tree.
// The front-end compiler will try to lookup by name the type and throw an exception if
// it's not defined. We can do fancy lazy stuff with this.
type Anonymous struct {
	Name    string
	Modules []string
}

func (_ *Anonymous) CandiceType() {}
func (a *Anonymous) String() string {
	if a == todoType {
		return ""
	}

	showcase := make([]string, len(a.Modules))
	for i, module := range a.Modules {
		showcase[i] = helper.RetrieveID(module)
	}
	return strings.Join(append(showcase, helper.RetrieveID(a.Name)), ".")
}
func (a *Anonymous) Alignment() int64 { return 0 }
func (a *Anonymous) SizeOf() int64    { return 0 }

type Struct struct {
	Fields []Type
	Names  []string
	Name   string
	ID     string
}

func (s *Struct) candiceType() {}

func (s *Struct) GetName() string {
	return s.Name
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
	str.WriteString("struct " + strings.Split(s.Name, "-")[0] + " {\n")
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
	return strings.Split(s.Name, "-")[0]
}

func (_ *Struct) CandiceType() {}

func (s *Struct) SizeOf() int64 {
	currentAddress := int64(0)
	// The padding formulas are quite confusing at first, check this out!
	// -- https://en.wikipedia.org/wiki/Data_structure_alignment#Computing_padding
	for _, t := range s.Fields {
		typeSize := t.SizeOf()
		alignment := t.Alignment()
		currentAddress += (currentAddress - (currentAddress % alignment)) % alignment
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

func IsFloat(t Type) bool {
	_, isFloat := t.(*Float)
	return isFloat
}

func IsPointer(t Type) bool {
	_, ok := t.(*Pointer)
	return ok
}

func IsArray(t Type) bool {
	_, ok := t.(*Array)
	return ok
}

func IsUnsignedInteger(t Type) bool {
	_, ok := t.(*UInteger)
	return ok
}

func IsFunction(t Type) bool {
	_, ok := t.(*Function)
	return ok
}

func UnwrapPossiblePointerAndDepth(t Type) (Type, int) {
	if pointer, isPointer := t.(*Pointer); isPointer {
		return pointer.UnwrapAndDepth()
	}

	return t, 0
}

func IsInteger(t Type) bool {
	_, ok := t.(*Integer)
	if !ok {
		_, ok = t.(*UInteger)
	}
	return ok
}

func IsUnion(t Type) bool {
	_, ok := t.(*Union)
	return ok
}
