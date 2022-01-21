package os

import (
  "fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

  "github.com/hofstadter-io/cuetils/utils"
)

type Stdout struct {}

func NewStdout(val cue.Value) (flow.Runner, error) {
  return &Stdout{}, nil
}

func (T* Stdout) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
	}

	v := t.Value()

	o := v.LookupPath(cue.ParsePath("#O"))

  fmt.Println(o)

	attr := v.Attribute("print")
	err = utils.PrintAttr(attr, v)

	return err
}
