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
  var fn string // filename
  var bs []byte // contents
  var m int64   // mode
  var a bool    // append

  ferr := func () error {
    ctx.CUELock.Lock()
    defer func() {
      ctx.CUELock.Unlock()
    }()
    var err error

    f := v.LookupPath(cue.ParsePath("filename"))

    fn, err = f.String()
    if err != nil {
      return err
    }

    // switch on c's type to fill appropriately
    c := v.LookupPath(cue.ParsePath("contents"))

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
    m, err = mode.Int64()
    if err != nil {
      return err
    }
    return nil

    av := v.LookupPath(cue.ParsePath("append"))
    a, err = av.Bool()
    if err != nil {
      return err
    }
    return nil
  }()
  if ferr != nil {
    return nil, ferr
  }

  // todo, check failure modes, fill, not return error?
  // (in all tasks)

  err := g_os.WriteFile(fn, bs, g_os.FileMode(m))
	return nil, err
}
