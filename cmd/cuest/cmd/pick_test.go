package cmd_test

import (
	"testing"

	"github.com/hofstadter-io/hof/lib/yagu"
	"github.com/hofstadter-io/hof/script/runtime"

	"github.com/hofstadter-io/cuetils/cmd/cuest/cmd"
)

func TestScriptPickCliTests(t *testing.T) {
	// setup some directories

	dir := "pick"

	workdir := ".workdir/cli/" + dir
	yagu.Mkdir(workdir)

	runtime.Run(t, runtime.Params{
		Setup: func(env *runtime.Env) error {
			// add any environment variables for your tests here

			return nil
		},
		Funcs: map[string]func(ts *runtime.Script, args []string) error{
			"__cuest": cmd.CallTS,
		},
		Dir:         "hls/cli/pick",
		WorkdirRoot: workdir,
	})
}