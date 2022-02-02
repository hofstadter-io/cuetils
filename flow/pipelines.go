package pipeline

import (
  "fmt"

  "cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
	"github.com/hofstadter-io/cuetils/pipeline/context"
	"github.com/hofstadter-io/cuetils/pipeline/pipe"
	"github.com/hofstadter-io/cuetils/structural"
	"github.com/hofstadter-io/cuetils/utils"
)

func hasPipelineAttr(val cue.Value, args []string) (attr cue.Attribute, found, keep bool) {
  attrs := val.Attributes(cue.ValueAttr)

  for _, attr := range attrs {
    if attr.Name() == "pipeline" {
      // found a pipeline, stop recursion
      found = true
      // if it matches our args, create and append
      keep = matchPipeline(attr, args)
      if keep {
        return attr, true, true
      }
    }
  }

  return cue.Attribute{}, found, false
}

func matchPipeline(attr cue.Attribute, args []string) (keep bool) {
  // fmt.Println("matching 1:", attr, args, len(args), attr.NumArgs())
  // if no args, match pipelines without args
  if len(args) == 0 {
    if attr.NumArgs() == 0 {
      return true
    }
    // extra check for one arg which is empty
    if attr.NumArgs() == 1 {
      s, err := attr.String(0)
      if err != nil {
        fmt.Println("bad pipeline tag:", err)
        return false
      }
      return s == ""
    }

    return false
  }

  // for now, match any
  // upgrade logic for user later
  for _, tag := range args {
    for p := 0; p < attr.NumArgs(); p++ {
      s, err := attr.String(p)
      if err != nil {
        fmt.Println("bad pipeline tag:", err)
        return false
      }
      if s == tag {
        return true
      }
    }
  }

  return false
}

func listPipelines(val cue.Value,  opts *flags.RootPflagpole, popts *flags.PipelineFlagpole) (error) {
  args := popts.Pipeline

  printer := func(v cue.Value) bool {
    attrs := v.Attributes(cue.ValueAttr)

    for _, attr := range attrs {
      if attr.Name() == "pipeline" {
        if len(args) == 0 || matchPipeline(attr, args) {
          if popts.Docs {
            s := ""
            docs := v.Doc()
            for _, d := range docs {
              s += d.Text()
            }
            fmt.Print(s)
          }
          if opts.Verbose {
            s, _ := utils.FormatCue(v)
            fmt.Printf("%s: %s\n", v.Path(), s)
          } else {
            fmt.Println(attr)
          }
        }
        return false
      }
    }

    return true
  }

  structural.Walk(val, printer, nil, walkOptions...)

  return nil
}

// maybe this becomes recursive so we can find
// nested pipelines, but not recurse when we find one
// only recurse when we have something which is not a pipeline or task
func findPipelines(ctx *context.Context, val cue.Value, opts *flags.RootPflagpole, popts *flags.PipelineFlagpole) ([]*pipe.Pipeline, error) {
  pipes := []*pipe.Pipeline{}

  // TODO, look for lists?
  s, err := val.Struct()
  if err != nil {
    // not a struct, so don't recurse
    // fmt.Println("not a struct", err)
    return nil, nil
  }

  args := popts.Pipeline

  // does our top-level (file-level) have @pipeline()
  _, found, keep := hasPipelineAttr(val, args)
  if keep  {
    // invoke TaskFactory
    p, err := pipe.NewPipeline(ctx, val)
    if err != nil {
      return pipes, err
    }
    pipes = append(pipes, p)
  }

  if found {
    return pipes, nil
  }

  iter := s.Fields(
		cue.Attributes(true),
		cue.Docs(true),
  )

  // loop over top-level fields in the cue instance
  for iter.Next() {
    v := iter.Value()

    _, found, keep := hasPipelineAttr(v, args)
    if keep  {
      p, err := pipe.NewPipeline(ctx, v)
      if err != nil {
        return pipes, err
      }
      pipes = append(pipes, p)
    }

    // break recursion if pipeline found
    if found {
      continue
    }

    // recurse to search for more pipelines
    ps, err := findPipelines(ctx, v, opts, popts)
    if err != nil {
      return pipes, nil 
    }
    pipes = append(pipes, ps...)
  }

  return pipes, nil
}

