package tasks

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"
)

// This function implements the Runner interface.
// It parses Cue values, you will see all of them recursively
func TaskFactory(P *Pipeline) func (cue.Value) (flow.Runner, error) {
  return func(val cue.Value) (flow.Runner, error) {

    // Check that we have something that looks like a task
    // (look for attributes that match cuetils ones)

    attrs := val.Attributes(cue.ValueAttr)
    // skip if no attributes
    if len(attrs) == 0 {
      return nil, nil
    }

    // fmt.Println("TF:", val.Path(), attrs)

    for _, attr := range attrs {
      // TODO, iterate over all attrs and build them up
      n := attr.Name()

      switch n {
      case "pipeline":
        t, err := maybePipeline(val, attr, P)
        if err != nil {
          fmt.Println("maybePipeline err:", err)
        }
        return t, err 
      case "task":
        t, err := maybeTask(val, attr, P)
        if err != nil {
          fmt.Println("maybeTask err:", err)
        }
        return t, err 
      }
    }

    return nil, nil
  }
}

func maybePipeline(val cue.Value, attr cue.Attribute, P *Pipeline) (flow.Runner, error) {
  // fmt.Println("maybePipeline:", attr)
  //fmt.Println(" -", P.Orig.Path(), val.Path())

  // how to know this is the root pipeline we are running?
  // if we return a Task for the root pipeline, we won't recurse
  //   ... or perhaps, not root, but "this"
  // it seems when we run a pipeline, we see the "this" value

  // right now, it seems we may be able to do this check
  if len(val.Path().Selectors()) == 0 {
    return nil, nil
  }

  // fmt.Println(" -- new'n")
  p, err := NewPipeline(val)
  if err != nil {
    return nil, err
  }

  np, _ := p.(*Pipeline)
  np.Prep()

  return p, nil
}

func maybeTask(val cue.Value, attr cue.Attribute, P *Pipeline) (flow.Runner, error) {
  // fmt.Println("maybeTask:", attr)
  if attr.NumArgs() == 0 {
    return nil, fmt.Errorf("No type provided to task: %s", attr)
  }

  taskId, err := attr.String(0)
  if err != nil {
    return nil, err
  }

  // fmt.Println("taskId:", taskId)

  // lookup task in task factory
  taskMaker, ok := TaskRegistry[taskId]
  if !ok {
    fmt.Println("uh oh") // this is not throwing an error, get here by having a bad task name
    return nil, fmt.Errorf("unknown task: %q", attr)
  }

  t, err := taskMaker(val)
  // fmt.Println("found task:", attr, t, err)

  return t, err 
}
