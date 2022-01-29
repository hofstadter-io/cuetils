package pipeline

import (
	"github.com/hofstadter-io/cuetils/pipeline/tasks"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/api"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/db"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/msg"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/os"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/pipe"
	"github.com/hofstadter-io/cuetils/pipeline/tasks/st"
)

func init() {
  tasks.TaskRegistry = tasks.TaskMap {
    "pipeline": pipe.NewPipeline,

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
    "os.Sleep": os.NewSleep,

    // api / db
    "api.Call": api.NewCall,
    "api.Serve": api.NewServe,
    "db.Call": db.NewQuery,

    // messaging
    "msg.IrcClient": msg.NewIrcClient,

    // websockets
    // message bus (rabbit,kafka,cloud,cloud-events)
    // channels / async

  }
}

