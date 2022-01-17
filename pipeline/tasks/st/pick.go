package st

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/structural"
	"github.com/hofstadter-io/cuetils/utils"
)

type PickTask struct {
	X   cue.Value `cue: "#X"`
	P   cue.Value `cue: "#P"`
	Ret cue.Value
}

// Tasks must implement a Run func, this is where we execute our task
func (P *PickTask) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
		// should we return?
	}

	v := t.Value()

	x := v.LookupPath(cue.ParsePath("#X"))
	p := v.LookupPath(cue.ParsePath("#P"))

	r, err := structural.PickValue(p, x, nil)
	if err != nil {
		return err
	}

	// Use fill to "return" a result to the workflow engine
	res := v.FillPath(cue.ParsePath("Out"), r)

	t.Fill(res)

	attr := v.Attribute("print")
	err = utils.PrintAttr(attr, res)

	return err
}
