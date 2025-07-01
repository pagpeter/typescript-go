package estransforms

import (
	"github.com/pagpeter/typescript-go/external/ast"
	"github.com/pagpeter/typescript-go/external/printer"
	"github.com/pagpeter/typescript-go/external/transformers"
)

type esDecoratorTransformer struct {
	transformers.Transformer
}

func (ch *esDecoratorTransformer) visit(node *ast.Node) *ast.Node {
	return node // !!!
}

func newESDecoratorTransformer(emitContext *printer.EmitContext) *transformers.Transformer {
	tx := &esDecoratorTransformer{}
	return tx.NewTransformer(tx.visit, emitContext)
}
