package compiler

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/microsoft/typescript-go/internal/core"
)

type concurrency struct {
	checkerCount int
}

func parseConcurrency(options *core.CompilerOptions, numFiles int) concurrency {
	if options.SingleThreaded.IsTrue() {
		return concurrency{
			checkerCount: 1,
		}
	}

	checkerCount := 4

	switch strings.ToLower(options.Concurrency) {
	case "default", "auto":
		break
	case "single", "none":
		checkerCount = 1
	case "max":
		checkerCount = runtime.GOMAXPROCS(0)
	case "half":
		checkerCount = max(1, runtime.GOMAXPROCS(0)/2)
	case "checker-per-file":
		checkerCount = -1
	default:
		if v, err := strconv.Atoi(options.Concurrency); err == nil && v > 0 {
			checkerCount = v
		}
	}

	return concurrency{
		checkerCount: checkerCount,
	}
}

func (c concurrency) isSingleThreaded() bool {
	return c.checkerCount == 1
}

func (c concurrency) getCheckerCount(numFiles int) int {
	if c.checkerCount == -1 {
		return max(1, numFiles)
	}
	return max(1, c.checkerCount)
}
