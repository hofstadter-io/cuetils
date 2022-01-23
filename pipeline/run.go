package pipeline

import (
	// "context"
	"fmt"
	"strings"
	// "time"

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
	if err != nil {
    s := structural.FormatCueError(err)
		return nil, fmt.Errorf("Error: %s", s)
	}
	if len(ins) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	// find  pipelines
  pipes := []*tasks.Pipeline{}
	for _, in := range ins {
    val := in.Value
    val, err = injectTags(val, popts.Tags)
    if err != nil {
      return nil, err
    }

    ps, err := findPipelines(val, opts, popts)
    if err != nil {
      return nil, err
    }
    pipes = append(pipes, ps...)
	}

  if len(pipes) == 0 {
    return nil, fmt.Errorf("no pipelines found")
  }

  // prep all of the pipelines and discover tasks
  for _, pipe := range pipes {
    err := pipe.Prep()
    if err != nil {
      return nil, err
    }
  }

  // start all of the pipelines
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

// maybe this becomes recursive so we can find
// nested pipelines, but not recurse when we find one
// only recurse when we have something which is not a pipeline or task
func findPipelines(val cue.Value, opts *flags.RootPflagpole, popts *flags.PipelineFlagpole) ([]*tasks.Pipeline, error) {
  pipes := []*tasks.Pipeline{}

  // TODO, look for lists?
  s, err := val.Struct()
  if err != nil {
    // not a struct, so don't recurse
    // fmt.Println("not a struct", err)
    return nil, nil
  }

  tags := popts.Pipeline

  // does our top-level (file-level) have @pipeline()
  _, found, keep := hasPipelineAttr(val, tags)
  if keep  {
    p, err := tasks.NewPipeline(val)
    if err != nil {
      return pipes, err
    }
    P, _ := p.(*tasks.Pipeline)
    pipes = append(pipes, P)
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

    _, found, keep := hasPipelineAttr(v, tags)
    if keep  {
      p, err := tasks.NewPipeline(v)
      if err != nil {
        return pipes, err
      }
      P, _ := p.(*tasks.Pipeline)
      pipes = append(pipes, P)
    }

    // break recursion if pipeline found
    if found {
      continue
    }

    // recurse to search for more pipelines
    ps, err := findPipelines(v, opts, popts)
    if err != nil {
      return pipes, nil 
    }
    pipes = append(pipes, ps...)
  }

  return pipes, nil
}

func hasPipelineAttr(val cue.Value, tags[]string) (attr cue.Attribute, found, keep bool) {
  attrs := val.Attributes(cue.ValueAttr)

  for _, attr := range attrs {
    if attr.Name() == "pipeline" {
      // found a pipeline, stop recursion
      found = true
      // if it matches our tags, create and append
      keep = matchPipeline(attr, tags)
      if keep  {
        return attr, true, true
      }
    }
  }

  return cue.Attribute{}, found, false
}

func matchPipeline(attr cue.Attribute, tags []string) (keep bool) {
  // fmt.Println("matching 1:", attr, tags, len(tags), attr.NumArgs())
  // if no tags, match pipelines without tags
  if len(tags) == 0 {
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
  for _, tag := range tags {
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

func injectTags(val cue.Value, tags []string) (cue.Value, error) {
  tagMap := make(map[string]string)
  for _, t := range tags {
    fs := strings.SplitN(t, "=", 2)
    if len(fs) != 2 {
      return val, fmt.Errorf("tags must have form key=value, got %q", t)
    }
    tagMap[fs[0]] =fs[1] 
  }
  return structural.InjectAttrsValue(val, "tag", tagMap)
}
