package st

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/structural"
	"github.com/hofstadter-io/cuetils/utils"
)

type Mask struct {}

func NewMask(val cue.Value) (flow.Runner, error) {
  return &Mask{}, nil
}

// Tasks must implement a Run func, this is where we execute our task
func (M *Mask) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
		// should we return?
	}

	v := t.Value()

	x := v.LookupPath(cue.ParsePath("val"))
	m := v.LookupPath(cue.ParsePath("mask"))

	r, err := structural.MaskValue(m, x, nil)
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
