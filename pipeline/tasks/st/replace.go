package st

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/structural"
	"github.com/hofstadter-io/cuetils/utils"
)

type Replace struct {}

func NewReplace(val cue.Value) (flow.Runner, error) {
  return &Replace{}, nil
}

// Tasks must implement a Run func, this is where we execute our task
func (M *Replace) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
		// should we return?
	}

	v := t.Value()

	x := v.LookupPath(cue.ParsePath("val"))
	repl := v.LookupPath(cue.ParsePath("repl"))

	r, err := structural.ReplaceValue(repl, x, nil)
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
