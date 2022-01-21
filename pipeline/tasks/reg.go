package tasks

import (
	"github.com/hofstadter-io/cuetils/pipeline/tasks/api"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/os"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/st"
)

type TaskMap map[string]TaskMaker

var TaskRegistry map[string]TaskMaker

func init() {
  TaskRegistry = TaskMap {
    "pipeline": NewPipeline,

    // structural
    "st/pick": st.NewPick,
    "st/mask": st.NewMask,
    "st/upsert": st.NewUpsert,

    // io
    // "os.Exec": os.NewExec,
    "os.Stdin": os.NewStdin,
    "os.Stdout": os.NewStdout,
    "os.ReadFile": os.NewReadFile,
    "os.WriteFile": os.NewWriteFile,

    // api / db
    "api/call": api.NewCall,

    // channels / async

  }
}

