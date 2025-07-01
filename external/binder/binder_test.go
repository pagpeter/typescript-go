package binder

import (
	"runtime"
	"testing"

	"github.com/pagpeter/typescript-go/external/ast"
	"github.com/pagpeter/typescript-go/external/core"
	"github.com/pagpeter/typescript-go/external/parser"
	"github.com/pagpeter/typescript-go/external/testutil/fixtures"
	"github.com/pagpeter/typescript-go/external/tspath"
	"github.com/pagpeter/typescript-go/external/vfs/osvfs"
)

func BenchmarkBind(b *testing.B) {
	for _, f := range fixtures.BenchFixtures {
		b.Run(f.Name(), func(b *testing.B) {
			f.SkipIfNotExist(b)

			fileName := tspath.GetNormalizedAbsolutePath(f.Path(), "/")
			path := tspath.ToPath(fileName, "/", osvfs.FS().UseCaseSensitiveFileNames())
			sourceText := f.ReadFile(b)

			compilerOptions := &core.CompilerOptions{Target: core.ScriptTargetESNext, Module: core.ModuleKindNodeNext}
			sourceAffecting := compilerOptions.SourceFileAffecting()

			parseOptions := ast.SourceFileParseOptions{
				FileName:         fileName,
				Path:             path,
				CompilerOptions:  sourceAffecting,
				JSDocParsingMode: ast.JSDocParsingModeParseAll,
			}
			scriptKind := core.GetScriptKindFromFileName(fileName)

			sourceFiles := make([]*ast.SourceFile, b.N)
			for i := range b.N {
				sourceFiles[i] = parser.ParseSourceFile(parseOptions, sourceText, scriptKind)
			}

			// The above parses do a lot of work; ensure GC is finished before we start collecting performance data.
			// GC must be called twice to allow things to settle.
			runtime.GC()
			runtime.GC()

			b.ResetTimer()
			for i := range b.N {
				BindSourceFile(sourceFiles[i])
			}
		})
	}
}
