//go:build !tsunsafe

package checker

import (
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
	_data       typeData // Type specific data
}

func (t *Type) data() typeData {
	return t._data
}

func (c *Checker) newType(flags TypeFlags, objectFlags ObjectFlags, data typeData) *Type {
	c.TypeCount++
	t := data.AsType()
	t.flags = flags
	t.objectFlags = objectFlags &^ (ObjectFlagsCouldContainTypeVariablesComputed | ObjectFlagsCouldContainTypeVariables | ObjectFlagsMembersResolved)
	t.id = TypeId(c.TypeCount)
	t.checker = c
	t._data = data
	return t
}
