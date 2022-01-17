package st

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/structural"
	"github.com/hofstadter-io/cuetils/utils"
)

type MaskTask struct {
	X   cue.Value `cue: "#X"`
	M   cue.Value `cue: "#M"`
	Ret cue.Value
}

// Tasks must implement a Run func, this is where we execute our task
func (M *MaskTask) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
		// should we return?
	}

	v := t.Value()

	x := v.LookupPath(cue.ParsePath("#X"))
	m := v.LookupPath(cue.ParsePath("#M"))

	r, err := structural.MaskValue(m, x, nil)
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
