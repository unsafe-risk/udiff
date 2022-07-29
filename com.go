package udiff

import "unsafe"

type opType uint16

type op struct {
	Type  opType
	Value unsafe.Pointer
}

func getValue[T any](o *op) *T {
	return (*T)(o.Value)
}

func setValue[T any](o *op, v *T) {
	o.Value = unsafe.Pointer(v)
}

type instruction struct {
	rtype uintptr
	code  []op
}
