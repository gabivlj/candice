package ctypes

import (
	"github.com/gabivlj/candice/pkg/a"
	"testing"
)

func TestStruct_Alignment(t *testing.T) {
	inner := &Struct{
		Fields: []Type{
			&Integer{BitSize: 32},
			&Integer{BitSize: 32},
			&Integer{BitSize: 64},
			&Integer{BitSize: 32},
		},
	}
	s := &Struct{
		Fields: []Type{
			&Integer{BitSize: 8},
			inner,
		},
	}
	a.Assert(s.SizeOf() == 32)
	a.Assert(s.Alignment() == 8)
	s = &Struct{
		Fields: []Type{
			&Array{Inner: &Integer{BitSize: 8}, Length: 1},
			&Array{Inner: &Integer{BitSize: 8}, Length: 15},
			&Integer{BitSize: 64},
			inner,
		},
	}
	a.Assert(s.SizeOf() == 48)
	a.Assert(s.Alignment() == 8)
}
