package os

import (
  "bufio"
  "fmt"

	"cuelang.org/go/cue"

  "github.com/hofstadter-io/cuetils/pipeline/context"
)

func init() {
  context.Register("os.Stdout", NewStdout)
}

type Stdout struct {}

func NewStdout(val cue.Value) (context.Runner, error) {
  return &Stdout{}, nil
}

func (T *Stdout) Run(ctx *context.Context) (interface{}, error) {
  bufStdout := bufio.NewWriter(ctx.Stdout)
  defer bufStdout.Flush()

  v := ctx.Value

  msg := v.LookupPath(cue.ParsePath("text")) 
  if msg.Err() != nil {
    return nil, msg.Err() 
  } else if msg.Exists() {
    // print strings directly to remove quotes
    if m, err := msg.String(); err == nil {
      fmt.Fprint(bufStdout, m)
    } else {
      fmt.Fprint(bufStdout, msg)
    }

  } else {
    err := fmt.Errorf("unknown msg:", msg)
    return nil, err
  }

	return nil, nil
}
