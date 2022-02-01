// This pipeline gets an api code with OAuth workflow
package twitch

import (
  "encoding/json"
  "github.com/hofstadter-io/cuetils/examples/streamer/auth"
)

vars: {
  title: string | *"" @tag(title)
  user: string | *"dr_verm" @tag(user)
}

meta: {
  @pipeline(meta)
  secrets: {
    env: { 
      TWITCH_CLIENT_ID: _ @task(os.Getenv)
    } 
    cid: env.TWITCH_CLIENT_ID

    tLoad: auth.load
    token: tLoad.token
  }

  twitch_req: {
    host: "https://api.twitch.tv"
    method: string | *"GET"
    headers: {
      "Client-ID": secrets.cid
      "Authorization": "Bearer \(secrets.token)"
    }
  }

}

user: {
  @pipeline(user)
  cfg: meta
  get: {
    username: _ | *vars.user
    @task(api.Call)
    req: cfg.twitch_req & {
      path: "/helix/users"
      query: {
        login: username
      }
    }
    resp: _
    user: resp.data[0]
  } 
  out: get.user
  str: json.Indent(json.Marshal(out), "", "  ")
  quiet: bool | *false
  if !quiet {
    print: { text: str + "\n" } @task(os.Stdout)
  }
}

channel: {
  @pipeline(channel)

  cfg: meta
  u: user & { 
    "cfg": cfg
    quiet: true
  }
  ug: u.out
  get: {
    @task(api.Call)
    req: cfg.twitch_req & {
      path: "/helix/channels"
      query: {
        broadcaster_id: ug.id
      }
    }
    resp: _
    channel: resp.data[0]
  }
  out: get.channel
  str: json.Indent(json.Marshal(out), "", "  ")
  quiet: bool | *false
  if !quiet {
    print: { text: str + "\n" } @task(os.Stdout)
  }
}

title: {
  @pipeline(title)

  cfg: meta
  u: user & { 
    "cfg": cfg
    quiet: true
  }
  ug: u.out

  // update stream title
  if vars.title != "" {
    debug: { text: "setting title to: '\(vars.title)'\n" } @task(os.Stdout)
    get: {
      @task(api.Call)
      req: cfg.twitch_req & {
        method: "PATCH"
        path: "/helix/channels"
        query: {
          broadcaster_id: ug.id
          title: vars.title
        }
      }
    }
    print: { text: get.resp } @task(os.Stdout)
  }
}
