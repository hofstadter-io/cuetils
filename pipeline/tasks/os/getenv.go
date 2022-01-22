package os

import (
	"fmt"
	g_os "os"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/utils"
)

type Getenv struct {}

func NewGetenv(val cue.Value) (flow.Runner, error) {
  return &Getenv{}, nil
}

func (T* Getenv) Run(t *flow.Task, err error) error {
	if err != nil {
		fmt.Println("Dep error", err)
	}

	v := t.Value()

  flds, err := v.Fields()
  if err != nil {
    return err
  }

  res := v
  for flds.Next() {
    sel := flds.Selector()
    key := sel.String()
    val := g_os.Getenv(key)
    res = v.FillPath(cue.MakePath(sel), val)
  }

	attr := v.Attribute("print")
	err = utils.PrintAttr(attr, v)

	// Use fill to "return" a result to the workflow engine
	t.Fill(res)

	return err
}
