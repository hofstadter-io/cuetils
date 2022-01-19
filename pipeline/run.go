package pipeline

import (
	"context"
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
	"github.com/hofstadter-io/cuetils/structural"
)

func Run(globs []string, opts *flags.RootPflagpole, popts *flags.PipelineFlagpole) ([]structural.GlobResult, error) {
	return run(globs, opts)
}

func run(globs []string, opts *flags.RootPflagpole) ([]structural.GlobResult, error) {
	ctx := cuecontext.New()

	vals, err := structural.ReadGlobs(globs, ctx, nil)
	if len(vals) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}
	if err != nil {
		return nil, err
	}

	// find tasks and then execute
	results := make([]structural.GlobResult, 0)
	for _, val := range vals {
		// walk tree, looking for `@pipeline(tags)`
		err = do(val, opts)
    if err != nil {
      return nil, err
    }
	}

	return results, nil
}

func do(in *structural.Input, opts *flags.RootPflagpole) error {
	var err error
	value := in.Value

	// Setup the flow Config
	cfg := &flow.Config{
		// make the Root, the Path to the value
		Root: cue.ParsePath("tasks"),
	}

	// fmt.Println("Dagging...")

	// create the workflow which will build the task graph
	workflow := flow.New(cfg, value, TaskFactory)

	// fmt.Println("Running...")

	// run our custom workflow
	err = workflow.Run(context.Background())
	if err != nil {
		s := structural.FormatCueError(err)
		return fmt.Errorf("Error: %s", s)
	}

	return nil
}
