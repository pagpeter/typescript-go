package estransforms

import (
	"github.com/pagpeter/typescript-go/external/ast"
	"github.com/pagpeter/typescript-go/external/printer"
	"github.com/pagpeter/typescript-go/external/transformers"
)

type exponentiationTransformer struct {
	transformers.Transformer
}

func (ch *exponentiationTransformer) visit(node *ast.Node) *ast.Node {
	return node // !!!
}

func newExponentiationTransformer(emitContext *printer.EmitContext) *transformers.Transformer {
	tx := &exponentiationTransformer{}
	return tx.NewTransformer(tx.visit, emitContext)
}
