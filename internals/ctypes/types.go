package ctypes

/// Candice types

type Type interface {
	// candiceType private flag
	candiceType()

	// SizeOf returns the size in bytes
	SizeOf() int64

	// Alignment returns the alignment in bytes
	Alignment() int64
}

type Void struct {}

func (_ *Void) candiceType(){}

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

func (_ *Integer) candiceType(){}

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

func (i *UInteger) SizeOf() int64 {
	return int64(i.BitSize / 8)
}

func (i *UInteger) Alignment() int64 {
	return int64(i.BitSize / 8)
}

func (_ *UInteger) candiceType(){}

type Pointer struct {
	Inner Type
}

func (_ *Pointer) Alignment() int64 {
	return 8
}

func (_ *Pointer) SizeOf() int64 { return 8 }

func (_ *Pointer) candiceType(){}

type Float struct {}

func (_ *Float) candiceType(){}

func (f *Float) SizeOf() int64 {
	return 8
}

func (_ *Float) Alignment() int64 {
	return 8
}

type Array struct {
	Inner Type
	Length int64
}

func (a *Array) SizeOf() int64 {
	return a.Inner.SizeOf() * a.Length
}

func (a *Array) Alignment() int64 {
	return a.Inner.Alignment()
}

func (_ *Array) candiceType(){}

type Struct struct {
	Fields []Type
	Names []string
}

func (_ *Struct) candiceType(){}

func (s *Struct) SizeOf() int64 {
	currentAddress := int64(0)
	// The padding formulas are quite confusing at first, check this out!
	// -- https://en.wikipedia.org/wiki/Data_structure_alignment#Computing_padding
	for _, t := range s.Fields  {
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
	max := int64(0)
	for _, t := range s.Fields {
		alignment := t.Alignment()
		if alignment >= max {
			max = alignment
		}
	}
	return max
}