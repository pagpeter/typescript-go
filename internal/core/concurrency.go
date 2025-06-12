package core

import (
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/microsoft/typescript-go/internal/testutil/race"
)

type Concurrency struct {
	checkerCount int
}

func ParseConcurrency(options *CompilerOptions) Concurrency {
	if options.SingleThreaded.IsTrue() {
		return Concurrency{
			checkerCount: 1,
		}
	}
	return parseConcurrency(options.Concurrency)
}

func parseConcurrency(v string) Concurrency {
	checkerCount := 4

	switch strings.ToLower(v) {
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
		if v, err := strconv.Atoi(v); err == nil && v > 0 {
			checkerCount = v
		}
	}

	return Concurrency{
		checkerCount: checkerCount,
	}
}

func (c Concurrency) SingleThreaded() bool {
	return c.checkerCount == 1
}

func (c Concurrency) CheckerCount(numFiles int) int {
	if c.checkerCount == -1 {
		return max(1, numFiles)
	}
	return max(1, c.checkerCount)
}

var testProgramConcurrency = sync.OnceValues(func() (concurrency Concurrency, raw string) {
	// Leave Program in SingleThreaded mode unless explicitly configured or in race mode.
	v := os.Getenv("TSGO_TEST_PROGRAM_CONCURRENCY")
	if v == "" && !race.Enabled {
		v = "single"
	}
	return parseConcurrency(v), v
})

func TestProgramConcurrency() (concurrency Concurrency, raw string) {
	return testProgramConcurrency()
}
