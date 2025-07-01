package ls

import (
	"github.com/pagpeter/typescript-go/external/compiler"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
)

type Host interface {
	GetProgram() *compiler.Program
	GetPositionEncoding() lsproto.PositionEncodingKind
	GetLineMap(fileName string) *LineMap
}
