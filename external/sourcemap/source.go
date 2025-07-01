package sourcemap

import "github.com/pagpeter/typescript-go/external/core"

type Source interface {
	Text() string
	FileName() string
	LineMap() []core.TextPos
}
