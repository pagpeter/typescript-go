package printer

import (
	"github.com/pagpeter/typescript-go/external/ast"
	"github.com/pagpeter/typescript-go/external/tspath"
)

type SourceFileMetaDataProvider interface {
	GetSourceFileMetaData(path tspath.Path) *ast.SourceFileMetaData
}
