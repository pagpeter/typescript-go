package project_test

import (
	"testing"

	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/project"
	"github.com/microsoft/typescript-go/internal/tspath"
	"gotest.tools/v3/assert"
)

func configFileExists(t *testing.T, service *project.Service, path tspath.Path, exists bool) {
	t.Helper()
	_, loaded := service.ConfigFileRegistry().ConfigFiles.Load(path)
	assert.Equal(t, loaded, exists, "config file %s should exist: %v", path, exists)
}

func serviceToPath(service *project.Service, fileName string) tspath.Path {
	return tspath.ToPath(fileName, service.GetCurrentDirectory(), service.FS().UseCaseSensitiveFileNames())
}

func stringifyReferences(references []string) string {
	var referencesToAdd []map[string]any
	for _, ref := range references {
		referencesToAdd = append(referencesToAdd, map[string]any{
			"path": ref,
		})
	}
	return core.Must(core.StringifyJson(referencesToAdd, "", "  "))
}

func ptrTo[T any](v T) *T {
	return &v
}
