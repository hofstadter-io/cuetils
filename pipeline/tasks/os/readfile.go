package os

import (
  "fmt"
  g_os "os"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

  "github.com/hofstadter-io/cuetils/utils"
)

type ReadFile struct {
  // might need F to be an object the helps us to
  // understand how to load the contents back in
  // such as a string, bytes, or a cue struct
  F cue.Value // the filename as a cue value 
  C cue.Value // the file contents
}

func NewReadFile(val cue.Value) (flow.Runner, error) {
  return &ReadFile{}, nil
}

func (T* ReadFile) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
	}

	v := t.Value()

	f := v.LookupPath(cue.ParsePath("f"))

  fn, err := f.String()
  if err != nil {
    return err
  }

  bs, err := g_os.ReadFile(fn)
  if err != nil {
    return err
  }

  // switch on c's type to fill appropriately
	c := v.LookupPath(cue.ParsePath("contents"))

  var res cue.Value
  switch k := c.IncompleteKind(); k {
  case cue.StringKind:
    res = v.FillPath(cue.ParsePath("contents"), string(bs))
  case cue.BytesKind:
    res = v.FillPath(cue.ParsePath("contents"), bs)

  case cue.StructKind:
    ctx := v.Context()
    c := ctx.CompileBytes(bs)
    if c.Err() != nil {
      return c.Err() 
    }
    res = v.FillPath(cue.ParsePath("contents"), c)

  default:
    return fmt.Errorf("Unsupported Content type in ReadFile task: %q", k)
  }


	// Use fill to "return" a result to the workflow engine
	t.Fill(res)

	attr := v.Attribute("print")
	err = utils.PrintAttr(attr, res)

	return err
}
