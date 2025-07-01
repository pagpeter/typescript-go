package estransforms

import (
	"github.com/pagpeter/typescript-go/external/ast"
	"github.com/pagpeter/typescript-go/external/printer"
	"github.com/pagpeter/typescript-go/external/transformers"
)

type logicalAssignmentTransformer struct {
	transformers.Transformer
}

func (ch *logicalAssignmentTransformer) visit(node *ast.Node) *ast.Node {
	return node // !!!
}

func newLogicalAssignmentTransformer(emitContext *printer.EmitContext) *transformers.Transformer {
	tx := &logicalAssignmentTransformer{}
	return tx.NewTransformer(tx.visit, emitContext)
}
