package os

import (
  "bufio"
  "fmt"
  "os"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

  "github.com/hofstadter-io/cuetils/utils"
)

type Stdout struct {}

func NewStdout(val cue.Value) (flow.Runner, error) {
  return &Stdout{}, nil
}

func (T* Stdout) Run(t *flow.Task, err error) error {
  bufStdout := bufio.NewWriter(os.Stdout)
  defer bufStdout.Flush()

	if err != nil {
		fmt.Println("Dep error", err)
	}

	v := t.Value()

  msg := v.LookupPath(cue.ParsePath("text")) 
  if msg.Err() != nil {
    return msg.Err() 
  } else if msg.Exists() {
    // print strings directly to remove quotes
    if m, err := msg.String(); err == nil {
      fmt.Fprint(bufStdout, m)
    } else {
      fmt.Fprint(bufStdout, msg)
    }

  } else {
    err := fmt.Errorf("unknown msg:", msg)
    return err
  }

	attr := v.Attribute("print")
	err = utils.PrintAttr(attr, v)

	return err
}
