package st

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/structural"
	"github.com/hofstadter-io/cuetils/utils"
)

type Pick struct {}

func NewPick(val cue.Value) (flow.Runner, error) {
  return &Pick{}, nil
}

// Tasks must implement a Run func, this is where we execute our task
func (P *Pick) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
		// should we return?
	}

	v := t.Value()

	x := v.LookupPath(cue.ParsePath("val"))
	p := v.LookupPath(cue.ParsePath("pick"))

	r, err := structural.PickValue(p, x, nil)
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
