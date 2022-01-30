// This pipeline gets an api code with OAuth workflow
package load

import (
  "encoding/json"

  "github.com/hofstadter-io/cuetils/examples/utils"
)

// twitch/auth/meta
meta: {
  @pipeline(meta,load) 

  vars: {
    RR: utils.RepoRoot
    root: RR.Out
    fn: "\(root)/examples/tmp/data.json"
  }
  secrets: {
    env: { 
      FOO: _ @task(os.Getenv)
    } 
    foo: env.FOO
  }
}

// twitch/auth/load
thing: {
  @pipeline(thing,load)

  cfg: meta

  files: { 
    t: { filename: cfg.vars.fn } @task(os.ReadFile)
    j: json.Unmarshal(t.contents)
  } 
  data: files.j
  say: data.cow

  debug: { text: "load/data: " + files.t.contents} @task(os.Stdout)
}

