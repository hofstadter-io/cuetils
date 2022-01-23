package tasks

import (
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/pipeline/tasks/api"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/db"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/os"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/st"
)

type TaskMap map[string]flow.TaskFunc

var TaskRegistry map[string]flow.TaskFunc

func init() {
  TaskRegistry = TaskMap {
    "pipeline": NewPipeline,

    "st.Diff": st.NewDiff,
    "st.Patch": st.NewPatch,
    "st.Pick": st.NewPick,
    "st.Mask": st.NewMask,
    "st.Insert": st.NewInsert,
    "st.Replace": st.NewReplace,
    "st.Upsert": st.NewUpsert,

    "os.Exec": os.NewExec,
    "os.Getenv": os.NewGetenv,
    "os.Stdin": os.NewStdin,
    "os.Stdout": os.NewStdout,
    "os.ReadFile": os.NewReadFile,
    "os.WriteFile": os.NewWriteFile,

    // api / db
    "api.Call": api.NewCall,
    "db.Call": db.NewQuery,

    // message bus (rabbit,kafka,cloud,cloud-events)

    // channels / async

  }
}

