//go:build tsunsafe

package ast

import (
	"sync/atomic"

	"github.com/microsoft/typescript-go/internal/core"
)

// AST Node
// Interface values stored in AST nodes are never typed nil values. Construction code must ensure that
// interface valued properties either store a true nil or a reference to a non-nil struct.

type Node struct {
	Kind   Kind
	Flags  NodeFlags
	Loc    core.TextRange
	id     atomic.Uint32
	Parent *Node
	_data  nodeData
}

func (n *Node) data() nodeData {
	return n._data
}

func newNode(kind Kind, data nodeData, hooks NodeFactoryHooks) *Node {
	n := data.AsNode()
	n.Loc = core.UndefinedTextRange()
	n.Kind = kind
	n._data = data
	if hooks.OnCreate != nil {
		hooks.OnCreate(n)
	}
	return n
}
