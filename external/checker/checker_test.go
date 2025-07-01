package checker_test

import (
	"testing"

	"github.com/pagpeter/typescript-go/external/ast"
	"github.com/pagpeter/typescript-go/external/bundled"
	"github.com/pagpeter/typescript-go/external/checker"
	"github.com/pagpeter/typescript-go/external/compiler"
	"github.com/pagpeter/typescript-go/external/core"
	"github.com/pagpeter/typescript-go/external/repo"
	"github.com/pagpeter/typescript-go/external/tsoptions"
	"github.com/pagpeter/typescript-go/external/tspath"
	"github.com/pagpeter/typescript-go/external/vfs/osvfs"
	"github.com/pagpeter/typescript-go/external/vfs/vfstest"
	"gotest.tools/v3/assert"
)

func TestGetSymbolAtLocation(t *testing.T) {
	t.Parallel()

	content := `interface Foo {
  bar: string;
}
declare const foo: Foo;
foo.bar;`
	fs := vfstest.FromMap(map[string]string{
		"/foo.ts": content,
		"/tsconfig.json": `
				{
					"compilerOptions": {},
					"files": ["foo.ts"]
				}
			`,
	}, false /*useCaseSensitiveFileNames*/)
	fs = bundled.WrapFS(fs)

	cd := "/"
	host := compiler.NewCompilerHost(nil, cd, fs, bundled.LibPath(), nil)

	parsed, errors := tsoptions.GetParsedCommandLineOfConfigFile("/tsconfig.json", &core.CompilerOptions{}, host, nil)
	assert.Equal(t, len(errors), 0, "Expected no errors in parsed command line")

	p := compiler.NewProgram(compiler.ProgramOptions{
		Config: parsed,
		Host:   host,
	})
	p.BindSourceFiles()
	c, done := p.GetTypeChecker(t.Context())
	defer done()
	file := p.GetSourceFile("/foo.ts")
	interfaceId := file.Statements.Nodes[0].Name()
	varId := file.Statements.Nodes[1].AsVariableStatement().DeclarationList.AsVariableDeclarationList().Declarations.Nodes[0].Name()
	propAccess := file.Statements.Nodes[2].AsExpressionStatement().Expression
	nodes := []*ast.Node{interfaceId, varId, propAccess}
	for _, node := range nodes {
		symbol := c.GetSymbolAtLocation(node)
		if symbol == nil {
			t.Fatalf("Expected symbol to be non-nil")
		}
	}
}

func TestCheckSrcCompiler(t *testing.T) {
	t.Parallel()

	repo.SkipIfNoTypeScriptSubmodule(t)
	fs := osvfs.FS()
	fs = bundled.WrapFS(fs)

	rootPath := tspath.CombinePaths(tspath.NormalizeSlashes(repo.TypeScriptSubmodulePath), "src", "compiler")

	host := compiler.NewCompilerHost(nil, rootPath, fs, bundled.LibPath(), nil)
	parsed, errors := tsoptions.GetParsedCommandLineOfConfigFile(tspath.CombinePaths(rootPath, "tsconfig.json"), &core.CompilerOptions{}, host, nil)
	assert.Equal(t, len(errors), 0, "Expected no errors in parsed command line")
	p := compiler.NewProgram(compiler.ProgramOptions{
		Config: parsed,
		Host:   host,
	})
	p.CheckSourceFiles(t.Context())
}

func BenchmarkNewChecker(b *testing.B) {
	repo.SkipIfNoTypeScriptSubmodule(b)
	fs := osvfs.FS()
	fs = bundled.WrapFS(fs)

	rootPath := tspath.CombinePaths(tspath.NormalizeSlashes(repo.TypeScriptSubmodulePath), "src", "compiler")

	host := compiler.NewCompilerHost(nil, rootPath, fs, bundled.LibPath(), nil)
	parsed, errors := tsoptions.GetParsedCommandLineOfConfigFile(tspath.CombinePaths(rootPath, "tsconfig.json"), &core.CompilerOptions{}, host, nil)
	assert.Equal(b, len(errors), 0, "Expected no errors in parsed command line")
	p := compiler.NewProgram(compiler.ProgramOptions{
		Config: parsed,
		Host:   host,
	})

	b.ReportAllocs()

	for b.Loop() {
		checker.NewChecker(p)
	}
}
