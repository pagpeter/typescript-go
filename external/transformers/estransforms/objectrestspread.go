package estransforms

import (
	"github.com/pagpeter/typescript-go/external/ast"
	"github.com/pagpeter/typescript-go/external/printer"
	"github.com/pagpeter/typescript-go/external/transformers"
)

type objectRestSpreadTransformer struct {
	transformers.Transformer
}

func (ch *objectRestSpreadTransformer) visit(node *ast.Node) *ast.Node {
	return node // !!!
}

func newObjectRestSpreadTransformer(emitContext *printer.EmitContext) *transformers.Transformer {
	tx := &objectRestSpreadTransformer{}
	return tx.NewTransformer(tx.visit, emitContext)
}
