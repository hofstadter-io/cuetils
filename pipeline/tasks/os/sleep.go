package os

import (
	"time"

	"cuelang.org/go/cue"

  "github.com/hofstadter-io/cuetils/pipeline/context"
)

func init() {
  context.Register("os.Sleep", NewSleep)
}

type Sleep struct {}

func NewSleep(val cue.Value) (context.Runner, error) {
  return &Sleep{}, nil
}

func (T *Sleep) Run(ctx *context.Context) (interface{}, error) {

	v := ctx.Value

  var D time.Duration
  d := v.LookupPath(cue.ParsePath("duration")) 
  if d.Err() != nil {
    return nil, d.Err()
  } else if d.Exists() {
    ds, err := d.String()
    if err != nil {
      return nil, err
    }
    D, err = time.ParseDuration(ds)
    if err != nil {
      return nil, err
    }
  }

  time.Sleep(D)
  res := v.FillPath(cue.ParsePath("done"), true)

	return res, nil
}

