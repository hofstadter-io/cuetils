package st

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/structural"
	"github.com/hofstadter-io/cuetils/utils"
)

type Diff struct {}

func NewDiff(val cue.Value) (flow.Runner, error) {
  return &Diff{}, nil
}

// Tasks must implement a Run func, this is where we execute our task
func (M *Diff) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
		// should we return?
	}

	v := t.Value()

	o := v.LookupPath(cue.ParsePath("orig"))
	n := v.LookupPath(cue.ParsePath("next"))

	r, err := structural.DiffValue(o, n, nil)
	if err != nil {
		return err
	}

	res := v.FillPath(cue.ParsePath("out"), r)

	attr := v.Attribute("print")
	err = utils.PrintAttr(attr, res)

	// Use fill to "return" a result to the workflow engine
	t.Fill(res)

	return err
}
