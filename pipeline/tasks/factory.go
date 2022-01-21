package tasks

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"
)

// This function implements the Runner interface.
// It parses Cue values, you will see all of them recursively
func TaskFactory(val cue.Value) (flow.Runner, error) {

	// Check that we have something that looks like a task
	// (look for attributes that match cuetils ones)

  attrs := val.Attributes(cue.ValueAttr)
  // skip if no attributes
  if len(attrs) == 0 {
    return nil, nil
  }

  for _, attr := range attrs {
    // TODO, iterate over all attrs and build them up
    n := attr.Name()

    switch n {
    case "pipeline":
      return maybePipeline(val, attr)
    case "task":
      return maybeTask(val, attr)
    }
  }

  return nil, nil
}

func maybePipeline(val cue.Value, attr cue.Attribute) (flow.Runner, error) {

  return nil, fmt.Errorf("error in maybePipeline")
}

func maybeTask(val cue.Value, attr cue.Attribute) (flow.Runner, error) {
  if attr.NumArgs() == 0 {
    return nil, fmt.Errorf("No type provided to task: %s", attr)
  }

  taskId, err := attr.String(0)
  if err != nil {
    return nil, err
  }

  // lookup task in task factory
  taskMaker, ok := TaskRegistry[taskId]
  if !ok {
    fmt.Errorf("unknown task: %q", attr)
  }

  return taskMaker(val)
}
