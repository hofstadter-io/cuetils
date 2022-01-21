package os

import (
  "fmt"
  g_os "os"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

  "github.com/hofstadter-io/cuetils/utils"
)

type WriteFile struct {
  // might need F to be an object the helps us to
  // understand how to load the contents back in
  // such as a string, bytes, or a cue struct
  F cue.Value // the filename as a cue value 
  C cue.Value // the file contents
}

func NewWriteFile(val cue.Value) (flow.Runner, error) {
  return &WriteFile{}, nil
}

func (T* WriteFile) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
	}

	v := t.Value()

	f := v.LookupPath(cue.ParsePath("filename"))

  fn, err := f.String()
  if err != nil {
    return err
  }

  // switch on c's type to fill appropriately
	c := v.LookupPath(cue.ParsePath("contents"))

  var bs []byte
  switch k := c.IncompleteKind(); k {
  case cue.StringKind:
    s, err := c.Bytes()
    if err != nil {
      return err
    }
    bs = []byte(s)
    
  case cue.BytesKind:
    bs, err = c.Bytes()
    if err != nil {
      return err
    }

  default:
    return fmt.Errorf("Unsupported content type in WriteFile task: %q", k)
  }

	mode := v.LookupPath(cue.ParsePath("mode"))
  m, err := mode.Int64()
  if err != nil {
    return err
  }

  err = g_os.WriteFile(fn, bs, g_os.FileMode(m))
  if err != nil {
    return err
  }

	attr := v.Attribute("print")
	err = utils.PrintAttr(attr, v)

	return err
}
