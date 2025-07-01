package ls

import (
	"context"

	"github.com/pagpeter/typescript-go/external/ast"
	"github.com/pagpeter/typescript-go/external/compiler"
	"github.com/pagpeter/typescript-go/external/lsp/lsproto"
)

type LanguageService struct {
	ctx        context.Context
	host       Host
	converters *Converters
}

func NewLanguageService(ctx context.Context, host Host) *LanguageService {
	return &LanguageService{
		host:       host,
		converters: NewConverters(host.GetPositionEncoding(), host.GetLineMap),
	}
}

// GetProgram updates the program if the project version has changed.
func (l *LanguageService) GetProgram() *compiler.Program {
	return l.host.GetProgram()
}

func (l *LanguageService) tryGetProgramAndFile(fileName string) (*compiler.Program, *ast.SourceFile) {
	program := l.GetProgram()
	file := program.GetSourceFile(fileName)
	return program, file
}

func (l *LanguageService) getProgramAndFile(documentURI lsproto.DocumentUri) (*compiler.Program, *ast.SourceFile) {
	fileName := DocumentURIToFileName(documentURI)
	program, file := l.tryGetProgramAndFile(fileName)
	if file == nil {
		panic("file not found: " + fileName)
	}
	return program, file
}
