package flow

import (
	go_ctx "context"
	"fmt"
	"os"
	"sync"

	// "time"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
	"github.com/hofstadter-io/cuetils/flow/context"
	"github.com/hofstadter-io/cuetils/flow/pipe"
	_ "github.com/hofstadter-io/cuetils/flow/tasks" // ensure tasks register
	"github.com/hofstadter-io/cuetils/structural"
	// "github.com/hofstadter-io/cuetils/utils"
)

/*
Input is to rigid
- reads from disk (can we workaround upstream?)
-
*/

func Run(globs []string, opts *flags.RootPflagpole, popts *flags.FlowFlagpole) ([]structural.GlobResult, error) {
	return run(globs, opts, popts)
}

// refactor out single/multi
func run(globs []string, opts *flags.RootPflagpole, popts *flags.FlowFlagpole) ([]structural.GlobResult, error) {
	ctx := cuecontext.New()

	ins, err := structural.ReadGlobs(globs, ctx, nil)
	if err != nil {
    s := structural.FormatCueError(err)
		return nil, fmt.Errorf("Error: %s", s)
	}
	if len(ins) == 0 {
		return nil, fmt.Errorf("no inputs found", '\n')
	}
    
  // sharedCtx := buildSharedContext

	// (refactor/pipe/many) find  flows
  pipes := []*pipe.Flow{}
	for _, in := range ins {

    // (refactor/pipe/solo)
    val := in.Value


    // (temp), give each own context (created in here), or maybe by flag? Need at least the shared mutex
    taskCtx, err := buildRootContext(val, opts, popts)
    // taskCtx, err := buildRootContext(sharedContex, val, opts, popts)
    if err != nil {
      return nil, err
    }

    // this might be buggy?
    val, err = injectTags(val, popts.Tags)
    if err != nil {
      return nil, err
    }

    // lets just print
    if popts.List {
      tags, secrets, errs := getTagsAndSecrets(val)
      if len(errs) > 0 {
        return nil, fmt.Errorf("in getTags: %v", errs)
      }
      if len(tags) > 0 {
        fmt.Println("tags:\n==============")
        for _, v := range tags {
          path := v.Path()
          fmt.Printf("%s: %# v %v\n", path, v, v.Attribute("tag"))
        }
        fmt.Println()
      }
      if len(secrets) > 0 {
        fmt.Println("secrets:\n==============")
        for _, v := range secrets {
          path := v.Path()
          fmt.Printf("%s: %# v %v\n", path, v, v.Attribute("secret"))
        }
        fmt.Println()
      }

      fmt.Println("flows:\n==============")
      err = listFlows(val, opts, popts)
      if err != nil {
        return nil, err
      } 

      continue
    }

    ps, err := findFlows(taskCtx, val, opts, popts)
    if err != nil {
      return nil, err
    }
    pipes = append(pipes, ps...)
	}

  if popts.List {
    return nil, nil
  }

  if len(pipes) == 0 {
    return nil, fmt.Errorf("no flows found")
  }

  // start all of the flows
  // TODO, use wait group, accume errors, flag for failure modes
  for _, pipe := range pipes {
    err := pipe.Start()
    if err != nil {
      return nil, err
    }
  }

  //time.Sleep(time.Second)
  //fmt.Println("done")
	return nil, nil
}


var walkOptions = []cue.Option{
  cue.Attributes(true),
  cue.Concrete(false),
  cue.Definitions(true),
  cue.Hidden(true),
  cue.Optional(true),
  cue.Docs(true),
}

func buildRootContext(val cue.Value, opts *flags.RootPflagpole, popts *flags.FlowFlagpole) (*context.Context, error) {
  // lookup the secret label in val
  // and build a filter write for stdout / stderr
  c := &context.Context{
    Stdin: os.Stdin,
    Stdout: os.Stdout,
    Stderr: os.Stderr,
    Context: go_ctx.Background(),
    CUELock: new(sync.Mutex),
    ValStore: new(sync.Map),
    Mailbox: new(sync.Map),
  }
  return c, nil
}
