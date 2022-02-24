package undomap

import "fmt"

type UndoMap[K comparable, T any] struct {
	values map[K][]T
	stack  []K
}

func (u *UndoMap[K, T]) String() string {
	return fmt.Sprintf("%v", u.values)
}

func New[K comparable, T any]() *UndoMap[K, T] {
	return &UndoMap[K, T]{values: map[K][]T{}, stack: []K{}}
}

func (u *UndoMap[K, T]) Add(key K, t T) {
	u.stack = append(u.stack, key)
	u.values[key] = append(u.values[key], t)
}

func (u *UndoMap[K, T]) Get(key K) T {
	var t T
	if values, ok := u.values[key]; !ok {
		return t
	} else {
		if len(values) == 0 {
			return t
		}
		return values[len(values)-1]
	}
}

func (u *UndoMap[K, T]) Pop() (K, T) {
	key := u.stack[len(u.stack)-1]
	u.stack = u.stack[:len(u.stack)-1]
	value := u.values[key][len(u.values[key])-1]
	u.values[key] = u.values[key][:len(u.values[key])-1]
	return key, value
}
