package use

import (
  "github.com/hofstadter-io/cuetils/pipeline/tests/repro/cuetils/load"
)

vars: {
  user: string | *"dr_verm" @tag(user)
}

// twitch/info/meta
meta: {
  @pipeline(meta,use)
  secrets: {
    env: { 
      COW: _ @task(os.Getenv)
    } 
    cow: env.COW

    tLoad: load.thing
    token: tLoad.say
  }

  req: {
    host: "https://postman-echo.com"
    method: "GET"
    query: {
      cow: "goes \(secrets.token)"
    }
  }

  debug: { text: "use/meta: " + secrets.token + "\n\n" } @task(os.Stdout)
}

// twitch/info/user
call: {
  @pipeline(call,use)

  cfg: meta

  get: {
    // @task(api.Call)
    req: cfg.req & {
      path: "/get"
      query: {
        login: vars.user
      }
    }
    resp: "dummy: \(cfg.req.method)"
    print: { text: resp } @task(os.Stdout)
  }
}
