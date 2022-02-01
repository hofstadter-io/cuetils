package os

import (
  "fmt"
  g_os "os"

	"cuelang.org/go/cue"

  "github.com/hofstadter-io/cuetils/pipeline/context"
)

func init() {
  context.Register("os.WriteFile", NewWriteFile)
}

type WriteFile struct {}

func NewWriteFile(val cue.Value) (context.Runner, error) {
  return &WriteFile{}, nil
}

func (T *WriteFile) Run(ctx *context.Context) (interface{}, error) {

	v := ctx.Value

	f := v.LookupPath(cue.ParsePath("filename"))

  fn, err := f.String()
  if err != nil {
    return nil, err
  }

  // switch on c's type to fill appropriately
	c := v.LookupPath(cue.ParsePath("contents"))

  var bs []byte
  switch k := c.IncompleteKind(); k {
  case cue.StringKind:
    s, err := c.Bytes()
    if err != nil {
      return nil, err
    }
    bs = []byte(s)
    
  case cue.BytesKind:
    bs, err = c.Bytes()
    if err != nil {
      return nil, err
    }

  default:
    return nil, fmt.Errorf("Unsupported content type in WriteFile task: %q", k)
  }

	mode := v.LookupPath(cue.ParsePath("mode"))
  m, err := mode.Int64()
  if err != nil {
    return nil, err
  }

  err = g_os.WriteFile(fn, bs, g_os.FileMode(m))
  if err != nil {
    return nil, err
  }

	return nil, err
}
