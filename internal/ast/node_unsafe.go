//go:build tsunsafe

package ast

import (
	"sync/atomic"
	"unsafe"

	"github.com/microsoft/typescript-go/internal/core"
)

// AST Node
// Interface values stored in AST nodes are never typed nil values. Construction code must ensure that
// interface valued properties either store a true nil or a reference to a non-nil struct.

type Node struct {
	Kind     Kind
	Flags    NodeFlags
	Loc      core.TextRange
	id       atomic.Uint32
	Parent   *Node
	dataType unsafe.Pointer
}

type interfaceValue struct {
	typ   unsafe.Pointer
	value unsafe.Pointer
}

func (n *Node) data() nodeData {
	var data nodeData
	(*interfaceValue)(unsafe.Pointer(&data)).typ = n.dataType
	(*interfaceValue)(unsafe.Pointer(&data)).value = unsafe.Pointer(n)
	return data
}

func newNode(kind Kind, data nodeData, hooks NodeFactoryHooks) *Node {
	n := data.AsNode()
	n.Loc = core.UndefinedTextRange()
	n.Kind = kind
	n.dataType = (*interfaceValue)(unsafe.Pointer(&data)).typ
	if hooks.OnCreate != nil {
		hooks.OnCreate(n)
	}
	return n
}
