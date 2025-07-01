package estransforms

import (
	"github.com/pagpeter/typescript-go/external/ast"
	"github.com/pagpeter/typescript-go/external/printer"
	"github.com/pagpeter/typescript-go/external/transformers"
)

type forawaitTransformer struct {
	transformers.Transformer
}

func (ch *forawaitTransformer) visit(node *ast.Node) *ast.Node {
	return node // !!!
}

func newforawaitTransformer(emitContext *printer.EmitContext) *transformers.Transformer {
	tx := &forawaitTransformer{}
	return tx.NewTransformer(tx.visit, emitContext)
}
