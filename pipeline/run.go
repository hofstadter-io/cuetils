package pipeline

import (
	// "context"
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
	"github.com/hofstadter-io/cuetils/pipeline/tasks"
	"github.com/hofstadter-io/cuetils/structural"
)

func Run(globs []string, opts *flags.RootPflagpole, popts *flags.PipelineFlagpole) ([]structural.GlobResult, error) {
	return run(globs, opts, popts)
}

func run(globs []string, opts *flags.RootPflagpole, popts *flags.PipelineFlagpole) ([]structural.GlobResult, error) {
	ctx := cuecontext.New()

	ins, err := structural.ReadGlobs(globs, ctx, nil)
	if len(ins) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}
	if err != nil {
		return nil, err
	}

  fmt.Println("num ins:", len(ins), popts.Tags)

	// find  pipelines
  pipes := []tasks.Pipeline{}
	for _, in := range ins {
    val := in.Value
    ps, err := findPipelines(val, opts, popts)
    if err != nil {
      return nil, err
    }
    pipes = append(pipes, ps...)
	}

  // loop over pipes, matching to flags?

	return nil, nil
}

func findPipelines(val cue.Value, opts *flags.RootPflagpole, popts *flags.PipelineFlagpole) ([]tasks.Pipeline, error) {
  // walk looking for pipelines
  // can we walk, or does cue/flow have to
  // in order for deps to work at the top level

  // filter for tags here

  return nil, nil
}

func do(in *structural.Input, opts *flags.RootPflagpole, popts *flags.PipelineFlagpole) error {
	// var err error


	// fmt.Println("Running...")

	// run our custom workflow
	//err = workflow.Run(context.Background())
	//if err != nil {
		//s := structural.FormatCueError(err)
		//return fmt.Errorf("Error: %s", s)
	//}

	return nil
}
