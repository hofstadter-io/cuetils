// This pipeline gets an api code with OAuth workflow

import (
  "encoding/json"
  "strings"

  "github.com/hofstadter-io/cuetils/examples/utils"
)

vars: {
  title: string | *"" @tag(title)
  user: string | *"dr_verm" @tag(user)
  user_id: string @tag(user_id)
}

meta: {
  @pipeline(meta)
  secrets: {
    env: { 
      TWITCH_CLIENT_ID: _ @task(os.Getenv)
    } 
    cid: env.TWITCH_CLIENT_ID

    r: utils.RepoRoot
    root: r.Out
    token_fn: "\(root)/examples/streamer/secrets/twitch.json"

    files: { 
      token_txt: { filename: token_fn } @task(os.ReadFile)
      token_json: json.Unmarshal(token_txt.contents)
    } 
    token: files.token_json.access_token
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
    @task(api.Call)
    req: cfg.twitch_req & {
      path: "/helix/users"
      query: {
        login: vars.user
      }
    }
  }

  id: get.resp.data[0].id
  print: { text: id + "\n" } @task(os.Stdout)
}

title: {
  @pipeline(title)

  cfg: meta

  // update stream title
  if vars.title != "" {
    debug: { text: "setting title to: '\(vars.title)'\n" } @task(os.Stdout)
    get: {
      @task(api.Call)
      req: cfg.twitch_req & {
        method: "PATCH"
        path: "/helix/channels"
        query: {
          broadcaster_id: vars.user_id
          title: vars.title
        }
      }
    }
    print: { text: get.resp } @task(os.Stdout)
  }

  // get and print current title
  if vars.title == "" {
    get: {
      @task(api.Call)
      req: cfg.twitch_req & {
        path: "/helix/channels"
        query: {
          broadcaster_id: vars.user_id
        }
      }
    }
    print: { text: get.resp } @task(os.Stdout)
  }
}
