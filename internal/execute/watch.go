package execute

import (
	"fmt"
	"time"

	"github.com/microsoft/typescript-go/internal/ast"
	"github.com/microsoft/typescript-go/internal/compiler"
	"github.com/microsoft/typescript-go/internal/incremental"
)

func start(w *watcher) ExitStatus {
	w.initialize()

	watchInterval := 1000 * time.Millisecond
	if w.options.ParsedConfig.WatchOptions != nil {
		watchInterval = time.Duration(*w.options.ParsedConfig.WatchOptions.Interval) * time.Millisecond
	}
	for {
		w.doCycle()
		time.Sleep(watchInterval)
	}
}

func (w *watcher) initialize() {
	// if this function is updated, make sure to update `StartForTest` in export_test.go as needed
	if w.configFileName == "" {
		w.host = compiler.NewCompilerHost(w.options.CompilerOptions(), w.sys.GetCurrentDirectory(), w.sys.FS(), w.sys.DefaultLibraryPath(), nil)
	}
	w.program = incremental.ReadBuildInfoProgram(w.options, incremental.NewBuildInfoReader(w.host))
}

func (w *watcher) doCycle() {
	// if this function is updated, make sure to update `RunWatchCycle` in export_test.go as needed

	if w.hasErrorsInTsConfig() {
		// these are unrecoverable errors--report them and do not build
		return
	}
	// updateProgram()
	w.program = incremental.NewProgram(compiler.NewProgram(compiler.ProgramOptions{
		Config:           w.options,
		Host:             w.host,
		JSDocParsingMode: ast.JSDocParsingModeParseForTypeErrors,
	}), w.program)

	if w.hasBeenModified(w.program.GetProgram()) {
		fmt.Fprint(w.sys.Writer(), "build starting at ", w.sys.Now(), w.sys.NewLine())
		timeStart := w.sys.Now()
		w.compileAndEmit()
		fmt.Fprint(w.sys.Writer(), "build finished in ", w.sys.Now().Sub(timeStart), w.sys.NewLine())
	} else {
		// print something???
		// fmt.Fprint(w.sys.Writer(), "no changes detected at ", w.sys.Now(), w.sys.NewLine())
	}
}
