// This pipeline gets an api code with OAuth workflow
package auth

import (
  "encoding/json"
  "strings"

  "github.com/hofstadter-io/cuetils/examples/utils"
)

// This pipeline will load a saved token
// for making calls to the Twitch API
load: {
  @pipeline(load,auth)

  cfg: meta

  files: { 
    token_txt: { filename: cfg.vars.token_fn } @task(os.ReadFile)
    token_json: json.Unmarshal(token_txt.contents)
  } 
  data: files.token_json
  token: data.access_token
}

meta: {
  @pipeline(meta,auth) 

  vars: {
    r: utils.RepoRoot
    root: r.Out
    code_fn: "\(root)/examples/streamer/secrets/twitch.code"
    token_fn: "\(root)/examples/streamer/secrets/twitch.json"
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
    callback: "http://localhost:2323/callback"
    domain: "https://id.twitch.tv"
    scopes: [
      "channel:manage:broadcast",
      "chat:read",
      "chat:edit",
    ]
    code_path: "/oauth2/authorize"
    code_query: {
      response_type: "code"
      state: "testing"
      scope: strings.Join(scopes," ")
      client_id: secrets.cid
      redirect_uri: callback 
      force_verify: "true"
    }
    cquery: [for k,v in code_query { "\(k)=\(v)" }]
    code_url: "\(domain)\(code_path)?" + strings.Join(cquery,"&")

    token_path: "/oauth2/token"
    token_query: {
      client_id: secrets.cid
      client_secret: secrets.key
      grant_type: "authorization_code"
      redirect_uri: callback 
    }
  }
}


// This pipeline will run the OAuth workflow
// go get a new token for the Twitch API
get_token: {
  @pipeline(get,auth)

  cfg: meta

  prompt: {
    @task(os.Stdout)
    text: """
    please open the following link in your browser

    \(cfg.twitch_cfg.code_url)

    you can ctrl-c this script after authorizing twitch
    """
  }

  server: {
    @pipeline(server,auth/get)

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
            html: """
            <html><body>
            <h1 style="margin-top:32px">
              OAuth token saved
            </h1>
            </body></html>
            """
          }
          // write auth code to file
          write_code: {
            @task(os.WriteFile)
            filename: "\(cfg.vars.code_fn)"
            contents: code
            mode: 0o666
          } 

          get_token: {
            @task(api.Call)
            req: {
              method: "POST"
              host: cfg.twitch_cfg.domain
              path: cfg.twitch_cfg.token_path
              query: cfg.twitch_cfg.token_query & {
                "code": code
              }
            }
            resp: string
          }

          write_token: {
            @task(os.WriteFile)
            filename: "\(cfg.vars.token_fn)"
            contents: get_token.resp
            mode: 0o666
          }
        }
      }
    }
  }
}

// This pipeline will refresh an existing token
// if it has not expired
refresh_token: {
  @pipeline(refresh)
}

