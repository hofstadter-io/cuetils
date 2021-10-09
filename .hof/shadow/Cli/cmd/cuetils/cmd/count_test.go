package cmd_test

import (
	"testing"

	"github.com/hofstadter-io/hof/lib/yagu"
	"github.com/hofstadter-io/hof/script/runtime"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/cmd"
)

func TestScriptCountCliTests(t *testing.T) {
	// setup some directories

	dir := "count"

	workdir := ".workdir/cli/" + dir
	yagu.Mkdir(workdir)

	runtime.Run(t, runtime.Params{
		Setup: func(env *runtime.Env) error {
			// add any environment variables for your tests here

			return nil
		},
		Funcs: map[string]func(ts *runtime.Script, args []string) error{
			"__cuetils": cmd.CallTS,
		},
		Dir:         "hls/cli/count",
		WorkdirRoot: workdir,
	})
}
