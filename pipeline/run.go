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

func Run(globs []string, opts *flags.RootPflagpole) ([]structural.GlobResult, error) {
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
		do(val, opts)
	}

	return results, nil
}

func do(in *structural.Input, opts *flags.RootPflagpole) error {
	var err error
	fmt.Println("Custom Flow Pipeline")

	fmt.Printf("%# v\n", in.Value)
	value := in.Value

	// Setup the flow Config
	cfg := &flow.Config{
		Root: cue.ParsePath("tasks"),
	}

	fmt.Println("Dagging...")

	// create the workflow which will build the task graph
	workflow := flow.New(cfg, value, TaskFactory)

	fmt.Println("Running...")

	// run our custom workflow
	err = workflow.Run(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
