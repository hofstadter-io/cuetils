package pipeline_test

import (
	"testing"

	"github.com/hofstadter-io/hof/lib/yagu"
	"github.com/hofstadter-io/hof/script/runtime"
)

func TestPipeline(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "write*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}
