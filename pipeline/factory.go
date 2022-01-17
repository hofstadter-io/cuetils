package pipeline

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/pipeline/tasks/api"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/st"
)

// This function implements the Runner interface.
// It parses Cue values, you will see all of them recursively
func TaskFactory(val cue.Value) (flow.Runner, error) {
	// You can see the recursive values with this
	// fmt.Println("TF: ", val)

	// Check that we have something that looks like a task
	// (look for attributes that match cuetils ones)
	attrs := val.Attributes(cue.ValueAttr)
	// fmt.Println("A:", attrs)

	if len(attrs) == 0 {
		return nil, nil
	}

	// should there every be more than 1 known attr?
	// probably not

	// look for an interesting attr in all the val's attrs
	// this will signify we have found a task
	for _, attr := range attrs {
		t, err := maybeTask(val, attr)
		if err != nil {
			return nil, err
		}
		// we found a task
		if t != nil {
			return t, nil
		}
	}

	// no attributes of interest
	return nil, nil
}

func maybeTask(val cue.Value, attr cue.Attribute) (flow.Runner, error) {
	n := attr.Name()
	if n != "task" {
		return nil, nil
	}

	if attr.NumArgs() == 0 {
		return nil, fmt.Errorf("No type provided to task:", attr)
	}

	t, err := attr.String(0)
	if err != nil {
		return nil, err
	}

	switch t {
	case "st/pick":
		return &st.PickTask{}, nil
	case "st/mask":
		return &st.MaskTask{}, nil
	case "st/upsert":
		return &st.UpsertTask{}, nil

	case "api/call":
		return &api.CallTask{}, nil

	default:
		fmt.Println("unknown attribute:", attr)
	}

	return nil, nil
}
