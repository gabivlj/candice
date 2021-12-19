package undomap

import "github.com/gabivlj/candice/internals/ctypes"

type UndoMap struct {
	values map[string][]ctypes.Type
	stack  []string
}

func New() *UndoMap {
	return &UndoMap{values: map[string][]ctypes.Type{}, stack: []string{}}
}

func (u *UndoMap) Add(key string, t ctypes.Type) {
	u.stack = append(u.stack, key)
	u.values[key] = append(u.values[key], t)
}

func (u *UndoMap) Get(key string) ctypes.Type {
	if values, ok := u.values[key]; !ok {
		return nil
	} else {
		if len(values) == 0 {
			return nil
		}
		return values[len(values)-1]
	}
}

func (u *UndoMap) Pop() (string, ctypes.Type) {
	key := u.stack[len(u.stack)-1]
	u.stack = u.stack[:len(u.stack)-1]
	value := u.values[key][len(u.values[key])-1]
	u.values[key] = u.values[key][:len(u.values[key])-1]
	return key, value
}
