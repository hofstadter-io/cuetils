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

  // If a struct, try to fill all fields with matching ENV VAR
  if v.IncompleteKind() == cue.StructKind {
    flds, err := v.Fields()
    if err != nil {
      return err
    }

    for flds.Next() {
      sel := flds.Selector()
      key := sel.String()
      val := g_os.Getenv(key)
      v = v.FillPath(cue.MakePath(sel), val)
    }
  } else {
    // otherwise, try to fill a field
    p := v.Path().Selectors()
    sel := p[len(p)-1]
    key := sel.String()
    val := g_os.Getenv(key)
    v = v.FillPath(cue.ParsePath(""), val)
  }


	attr := v.Attribute("print")
	err = utils.PrintAttr(attr, v)

	// Use fill to "return" a result to the workflow engine
	t.Fill(v)

	return err
}
