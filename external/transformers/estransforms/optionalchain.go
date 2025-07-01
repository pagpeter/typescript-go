package estransforms

import (
	"github.com/pagpeter/typescript-go/external/ast"
	"github.com/pagpeter/typescript-go/external/printer"
	"github.com/pagpeter/typescript-go/external/transformers"
)

type optionalChainTransformer struct {
	transformers.Transformer
}

func (ch *optionalChainTransformer) visit(node *ast.Node) *ast.Node {
	return node // !!!
}

func newOptionalChainTransformer(emitContext *printer.EmitContext) *transformers.Transformer {
	tx := &optionalChainTransformer{}
	return tx.NewTransformer(tx.visit, emitContext)
}
