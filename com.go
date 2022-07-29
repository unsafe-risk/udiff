package udiff

import (
	"reflect"
	"sync"
	"unsafe"
)

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

var buildTable sync.Map

func build(t reflect.Type) *instruction {
	if v, ok := buildTable.Load(t); ok {
		return v.(*instruction)
	}
	inst := &instruction{}
	buildTable.Store(t, inst) // Early Store to prevent recursive calls
	switch t.Kind() {

	}
	return nil
}
