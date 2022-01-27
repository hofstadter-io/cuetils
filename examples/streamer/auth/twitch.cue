// This pipeline tests OAuth workflow
// with the twitch API

import (
  "encoding/json"
  "strings"

  "github.com/hofstadter-io/cuetils/examples/utils"
)

@pipeline()

vars: {
  r: utils.RepoRoot
  root: r.Out
  code_fn: "\(vars.root)/examples/streamer/secrets/twitch.code"
}

secrets: {
  env: { 
    TWITCH_CLIENT_ID: _ @task(os.Getenv)
    TWITCH_SECRET_KEY: _ @task(os.Getenv)
  } 
  cid: env.TWITCH_CLIENT_ID
  key: env.TWITCH_SECRET_KEY
}

twitch_cfg: {
  auth_domain: "https://id.twitch.tv"
  auth_path: "/oauth2/authorize"
  auth_url: "\(auth_domain)\(auth_path)"
  scopes: [
    "user:read:email",
  ]
  auth_query: {
    response_type: "code"
    redirect_uri: "http://localhost:2323/callback"
    state: "testing"
    scope: strings.Join(scopes,",")
    client_id: secrets.cid
  }
  query: [for k,v in auth_query { "\(k)=\(v)" }]

  url: "\(auth_url)?" + strings.Join(query,"&")
}

prompt: {
  @task(os.Stdout)
  text: """
  please open the following link in your browser

  \(twitch_cfg.url)

  you can ctrl-c this script after authorizing twitch

  """
}

server: {
  @pipeline(server)

  run: {
    @task(api.Serve)
    port: "2323"
    routes: {
      "/callback": {
        @pipeline()
        req: _
        code: req.query.code[0]
        resp: {
          status: 200
          body: "code received"
        }
        // write auth code to file
        write: {
          @task(os.WriteFile)
          filename: "\(vars.code_fn)"
          contents: code
          mode: 0o666
        } 
      }
    }
  }
}
