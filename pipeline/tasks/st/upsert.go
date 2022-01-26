package st

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/structural"
	"github.com/hofstadter-io/cuetils/utils"
)

type Upsert struct {}

func NewUpsert(val cue.Value) (flow.Runner, error) {
  return &Upsert{}, nil
}

// Tasks must implement a Run func, this is where we execute our task
func (U *Upsert) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
		// should we return?
	}

	// not sure this is OK, but the value which was used for this task
	v := t.Value()
	x := v.LookupPath(cue.ParsePath("val"))
	u := v.LookupPath(cue.ParsePath("up"))

	r, err := structural.UpsertValue(u, x, nil)
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