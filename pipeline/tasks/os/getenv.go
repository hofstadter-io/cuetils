package os

import (
	g_os "os"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/utils"
)

type Getenv struct {
  Orig cue.Value
}

func NewGetenv(val cue.Value) (flow.Runner, error) {
  return &Getenv{ Orig: val }, nil
}

func (T* Getenv) Run(t *flow.Task, err error) error {
	// v := T.Orig
  v := t.Value()

  // fmt.Println("getenv(deps):", t.Dependencies())

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
      // fmt.Println("pdeps:", t.PathDependencies(cue.MakePath(sel)))
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
