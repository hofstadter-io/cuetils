import "strings"

IRCHandler: {
  // input
  msg: {
    Command: "PRIVMSG"
    Params: ["#dr_verm", "!today"]
  }

  // output (oneof), first taken anyhow
  error?: _
  resp?: _
  pipe?: _

  if msg.Command == "PRIVMSG" {
    cmd: msg.Params[1]

    switch: [
      if simpleCmds[cmd] != _|_ {
        resp: simpleCmds[cmd]
      }
      if cmd == "!docker" { pipe: IRCHandlers.docker },
      if strings.HasPrefix(cmd, "!so") { pipe: IRCHandlers.shoutout & { args: cmd } },

      { error: "unknown cmd: " + cmd },
    ] 

    switch[0]
  }

  if msg.Command == "USEREVENT" {

  }
}

simpleCmds: {
  "!today": "working on a twitch bot in CUEtils pipelines"
  "!music": "streaming https://www.youtube.com/watch?v=udGvUx70Q3U"
}

IRCHandlers: {
  docker: {
    @pipeline()

    args?: _

    get: {
      @task(os.Exec)
      cmd: ["docker", "ps"]
      stdout: string  
    }
  
    resp: get.stdout
  }

  shoutout: {
    @pipeline()

    args: string
    who: strings.TrimPrefix(args, "!so @")

    get: {
      // call twitch api for info about the user
      // eventually also look up custom data in DB
    }

    resp: "not implemented yet"

  }
}
