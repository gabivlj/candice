package undomap

import (
	"testing"

	"github.com/gabivlj/candice/internals/ctypes"
	"github.com/gabivlj/candice/pkg/a"
)

func TestUndoMap_AddPop(t *testing.T) {
	undoMap := New[string, ctypes.Type]()
	undoMap.Add("integerValue", &ctypes.Integer{BitSize: 32})
	undoMap.Add("integerValue", &ctypes.Integer{BitSize: 33})
	undoMap.Add("integerValue", &ctypes.Integer{BitSize: 34})
	undoMap.Add("integerValue", &ctypes.Integer{BitSize: 35})
	key, value := undoMap.Pop()
	a.AssertEqual(key, "integerValue")
	a.Assert(value.(*ctypes.Integer).BitSize == 35, "35")
	key, value = undoMap.Pop()
	a.AssertEqual(key, "integerValue")
	a.Assert(value.(*ctypes.Integer).BitSize == 34, "34")
	key, value = undoMap.Pop()
	a.AssertEqual(key, "integerValue")
	a.Assert(value.(*ctypes.Integer).BitSize == 33, "33")
	key, value = undoMap.Pop()
	a.AssertEqual(key, "integerValue")
	a.Assert(value.(*ctypes.Integer).BitSize == 32, "32")
	a.Assert(undoMap.Get("integerValue") == nil)
}
