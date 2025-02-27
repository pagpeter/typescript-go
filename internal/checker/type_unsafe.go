//go:build tsunsafe

package checker

import (
	"unsafe"

	"github.com/microsoft/typescript-go/internal/ast"
)

// Type

type Type struct {
	flags       TypeFlags
	objectFlags ObjectFlags
	id          TypeId
	symbol      *ast.Symbol
	alias       *TypeAlias
	checker     *Checker
	dataType    unsafe.Pointer // Type specific data
}

type interfaceValue struct {
	typ   unsafe.Pointer
	value unsafe.Pointer
}

func (t *Type) data() typeData {
	var data typeData
	(*interfaceValue)(unsafe.Pointer(&data)).typ = t.dataType
	(*interfaceValue)(unsafe.Pointer(&data)).value = unsafe.Pointer(t)
	return data
}

func (c *Checker) newType(flags TypeFlags, objectFlags ObjectFlags, data typeData) *Type {
	c.TypeCount++
	t := data.AsType()
	t.flags = flags
	t.objectFlags = objectFlags &^ (ObjectFlagsCouldContainTypeVariablesComputed | ObjectFlagsCouldContainTypeVariables | ObjectFlagsMembersResolved)
	t.id = TypeId(c.TypeCount)
	t.checker = c
	t.dataType = (*interfaceValue)(unsafe.Pointer(&data)).typ
	return t
}
