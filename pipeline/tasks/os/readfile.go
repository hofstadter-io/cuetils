package os

import (
  "fmt"
  g_os "os"

	"cuelang.org/go/cue"

  "github.com/hofstadter-io/cuetils/pipeline/context"
)

func init() {
  context.Register("os.ReadFile", NewReadFile)
}

type ReadFile struct {}

func NewReadFile(val cue.Value) (context.Runner, error) {
  return &ReadFile{}, nil
}

func (T *ReadFile) Run(ctx *context.Context) (interface{}, error) {

	v := ctx.Value

	f := v.LookupPath(cue.ParsePath("filename"))

  fn, err := f.String()
  if err != nil {
    return nil, err
  }

  bs, err := g_os.ReadFile(fn)
  if err != nil {
    return nil, err
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
      return nil, c.Err() 
    }
    res = v.FillPath(cue.ParsePath("contents"), c)

  case cue.BottomKind:
    res = v.FillPath(cue.ParsePath("contents"), string(bs))

  default:
    return nil, fmt.Errorf("Unsupported Content type in ReadFile task: %q", k)
  }

	return res, nil
}
