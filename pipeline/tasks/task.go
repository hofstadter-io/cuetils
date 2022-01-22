package tasks

import (
	"context"
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/structural"
  "github.com/hofstadter-io/cuetils/utils"
)

type Pipeline struct {
  Orig cue.Value

  Ctrl *flow.Controller
}

func NewPipeline(val cue.Value) (flow.Runner, error) {
  p := &Pipeline{
    Orig: val,
  }
  return p, nil
}

func (P *Pipeline) Prep() error {
  // fmt.Println("prepping:", P.Orig.Path(), P.Orig.Attributes(cue.ValueAttr))

	// Setup the flow Config
	cfg := &flow.Config{}

  v := P.Orig.Context().CompileString("{...}")
  u := v.Unify(P.Orig) 

  //s, _ := utils.PrintCue(u)
  //fmt.Printf("===\n%s\n===\n", s)

	// create the workflow which will build the task graph
	P.Ctrl = flow.New(cfg, u, TaskFactory(P))

  return nil
}

func (P *Pipeline) Start() error {
  // fmt.Println("beg: pipe", P.Orig.Attributes(cue.ValueAttr))
  // fmt.Println("starting:", P.Orig.Attributes(cue.ValueAttr))
  _, err := P.Ctrl.Run(context.Background())
  // fmt.Println("finishing:", P.Orig.Attributes(cue.ValueAttr))

  if err != nil {
    s := structural.FormatCueError(err)
		return fmt.Errorf("Error: %s", s)
  }
  // fmt.Println("end: pipe", P.Orig.Attributes(cue.ValueAttr))
  return nil
}

// for recursively included pipelines?
func (P *Pipeline) Run(t *flow.Task, err error) error {
  // fmt.Println("beg: pipe", P.Orig.Attributes(cue.ValueAttr))
	if err != nil {
		fmt.Println("Dep error", err)
		// should we return?
	}

  // run the pipeline
  final, err := P.Ctrl.Run(context.Background())
  // did it error?
  if err != nil {
    s := structural.FormatCueError(err)
		return fmt.Errorf("Error: %s", s)
  }

  attr := final.Attribute("print")
  err = utils.PrintAttr(attr, final)

	t.Fill(final)

  // fmt.Println("end: pipe", P.Orig.Attributes(cue.ValueAttr))
	return err
}
