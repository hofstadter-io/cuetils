package os

import (
	"fmt"
	"time"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/utils"
)

type Sleep struct {}

func NewSleep(val cue.Value) (flow.Runner, error) {
  return &Sleep{}, nil
}

func (T* Sleep) Run(t *flow.Task, err error) error {
	if err != nil {
		fmt.Println("Dep error", err)
	}

	v := t.Value()

  d := v.LookupPath(cue.ParsePath("duration")) 
  if d.Err() != nil {
    return err
  } else if d.Exists() {
    ds, err := d.String()
    if err != nil {
      return err
    }
    D, err := time.ParseDuration(ds)
    if err != nil {
      return err
    }

    time.Sleep(D)
  }

  res := v.FillPath(cue.ParsePath("done"), true)
	attr := v.Attribute("print")
	err = utils.PrintAttr(attr, v)

	// Use fill to "return" a result to the workflow engine
	t.Fill(res)

	return err
}

