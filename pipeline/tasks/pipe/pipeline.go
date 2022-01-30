package pipe

import (
	"context"
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

  "github.com/hofstadter-io/cuetils/pipeline/tasks"
	"github.com/hofstadter-io/cuetils/structural"
  "github.com/hofstadter-io/cuetils/utils"
)

type Pipeline struct {
  Orig cue.Value
  Final cue.Value

  path string
  rpath string

  Ctrl *flow.Controller
}

func NewPipeline(val cue.Value) (flow.Runner, error) {
  _, rp := val.ReferencePath()
  p := &Pipeline{
    Orig: val,
    path: val.Path().String(),
    rpath: rp.String(),
  }
  return p, nil
}

func (P *Pipeline) prep(val cue.Value) error {
	// Setup the flow Config
	cfg := &flow.Config{
		// InferTasks:     true,
		IgnoreConcrete: true,
    UpdateFunc: func(c *flow.Controller, t *flow.Task) error {
      //if t != nil {
        //fmt.Println("  UPDATE:", P.path, t.Index(), t.Path(), t.State())
        //// P.Final = t.Value()
      //} else {
        //fmt.Println("  UPDATE:", P.path, "<nil task>")
      //}
      return nil
    },
  }

  v := P.Orig.Context().CompileString("{...}")
  u := v.Unify(val) 

	// create the workflow which will build the task graph
	P.Ctrl = flow.New(cfg, u, tasks.TaskFactory())

  return nil
}

func (P *Pipeline) run(val cue.Value) error {
  // fmt.Println("pipe(beg):", P.path, P.rpath)
  P.prep(val)

  //tasks := P.Ctrl.Tasks()
  //fmt.Println("  tasks:")
  //for _,t := range tasks {
    //fmt.Println("  -", t.Index(), t.Path(), t.State())
    //for _, d := range t.Dependencies() {
      //fmt.Println("    -", d.Index(), d.Path(), d.State())
    //}
  //}

  final, err := P.Ctrl.Run(context.Background())

  // fmt.Println("pipe(end):", P.path, P.rpath)
  P.Final = final
  if err != nil {
    s := structural.FormatCueError(err)
		return fmt.Errorf("Error: %s", s)
  }

  return nil
}

// This is for top-level pipelines
func (P *Pipeline) Start() error {
  return P.run(P.Orig)
}

// This is for included pipelines or nested under other pipelines
func (P *Pipeline) Run(t *flow.Task, err error) error {
	if err != nil {
		fmt.Println("Dep error", err)
		// should we return?
	}

  err = P.run(t.Value())
  // err = P.run(P.Orig)
  if err != nil {
    return err
  }

  attr := P.Final.Attribute("print")
  err = utils.PrintAttr(attr, P.Final)

	t.Fill(P.Final)

	return err
}
