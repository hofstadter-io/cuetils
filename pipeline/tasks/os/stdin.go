package os

import (
  "bufio"
  "fmt"
  g_os "os"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

  "github.com/hofstadter-io/cuetils/utils"
)

type Stdin struct {}

func NewStdin(val cue.Value) (flow.Runner, error) {
  return &Stdin{}, nil
}

func (T* Stdin) Run(t *flow.Task, err error) error {

	if err != nil {
		fmt.Println("Dep error", err)
	}

	v := t.Value()

  msg := v.LookupPath(cue.ParsePath("msg")) 
  if msg.Err() != nil {
    return err
  } else if msg.Exists() {
    m, err := msg.String()
    if err != nil {
      return err
    }
    // print message to user
    fmt.Print(m)
  }

  reader := bufio.NewReader(g_os.Stdin)
  text, _ := reader.ReadString('\n')

  res := v.FillPath(cue.ParsePath("contents"), text)
	attr := v.Attribute("print")
	err = utils.PrintAttr(attr, v)

	// Use fill to "return" a result to the workflow engine
	t.Fill(res)

	return err
}
