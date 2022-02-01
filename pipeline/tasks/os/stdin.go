package os

import (
  "bufio"
  "fmt"

	"cuelang.org/go/cue"

  "github.com/hofstadter-io/cuetils/pipeline/context"
)

func init() {
  context.Register("os.Stdin", NewStdin)
}

type Stdin struct {}

func NewStdin(val cue.Value) (context.Runner, error) {
  return &Stdin{}, nil
}

func (T *Stdin) Run(ctx *context.Context) (interface{}, error) {

	v := ctx.Value

  msg := v.LookupPath(cue.ParsePath("msg")) 
  if msg.Err() != nil {
    return nil, msg.Err()

  } else if msg.Exists() {
    m, err := msg.String()
    if err != nil {
      return nil, err
    }
    // print message to user
    fmt.Fprint(ctx.Stdout, m)
  }

  reader := bufio.NewReader(ctx.Stdin)
  text, _ := reader.ReadString('\n')

  res := v.FillPath(cue.ParsePath("contents"), text)

	return res, nil 
}
