package tasks

import (
  "fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"
)

type TaskMaker func (cue.Value) (flow.Runner, error)

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

func (P *Pipeline) Prep(val cue.Value) {
	// Setup the flow Config
	cfg := &flow.Config{}

	// create the workflow which will build the task graph
	P.Ctrl = flow.New(cfg, val, TaskFactory)
}

func (P *Pipeline) Run(t *flow.Task, err error) error {
	if err != nil {
		fmt.Println("Dep error", err)
		// should we return?
	}

	// v := t.Value()

	//in := v.LookupPath(cue.ParsePath("In"))
	//ts := v.LookupPath(cue.ParsePath("Tasks"))

	//// Use fill to "return" a result to the workflow engine
	//res := v.FillPath(cue.ParsePath("Out"), r)

	//t.Fill(res)

	//attr := v.Attribute("print")
	//err = utils.PrintAttr(attr, res)

	return err
}
