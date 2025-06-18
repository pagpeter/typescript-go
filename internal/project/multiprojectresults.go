package project

import (
	"context"

	"github.com/microsoft/typescript-go/internal/collections"
	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/ls"
	"github.com/microsoft/typescript-go/internal/lsp/lsproto"
	"github.com/microsoft/typescript-go/internal/tspath"
)

func GetPerProjectResults[T any](
	service *Service,
	url lsproto.DocumentUri,
	ctx context.Context,
	requestToResult func(info *ScriptInfo, project *Project, ctx context.Context, filePathForRequest tspath.Path) []T,
	onResult func(result T),
) error {
	defaultInfo, defaultProject, projects := service.getProjectsForURI(url)
	var workInProgress collections.SyncMap[*Project, bool]
	results := map[*Project][]T{}
	fileName := ls.DocumentURIToFileName(url)
	filePathOfRequest := service.toPath(fileName)
	wg := core.NewWorkGroup(false)
	workInProgress.Store(defaultProject, true)
	wg.Queue(func() {
		results[defaultProject] = requestToResult(
			defaultInfo,
			defaultProject,
			ctx,
			filePathOfRequest,
		)
		// TODO:: if default location needs to be add more projects to request
	})
	for info, projectsForInfo := range projects {
		for project := range projectsForInfo.Keys() {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			if _, loaded := workInProgress.LoadOrStore(project, true); !loaded {
				continue
			}
			wg.Queue(func() {
				results[project] = requestToResult(
					info,
					project,
					ctx,
					filePathOfRequest,
				)
			})
		}
	}
	wg.RunAndWait()

	for _, result := range results[defaultProject] {
		onResult(result)
	}
	for project, results := range results {
		if project != defaultProject {
			for _, result := range results {
				onResult(result)
			}
		}
	}
	return nil
}
