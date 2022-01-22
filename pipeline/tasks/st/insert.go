package st

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/structural"
	"github.com/hofstadter-io/cuetils/utils"
)

type Insert struct {}

func NewInsert(val cue.Value) (flow.Runner, error) {
  return &Insert{}, nil
}

// Tasks must implement a Run func, this is where we execute our task
func (M *Insert) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
		// should we return?
	}

	v := t.Value()

	x := v.LookupPath(cue.ParsePath("val"))
	ins := v.LookupPath(cue.ParsePath("ins"))

	r, err := structural.InsertValue(ins, x, nil)
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
