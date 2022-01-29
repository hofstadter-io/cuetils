// This pipeline listens to the channel chat
// and does things, hopefully

import "strings"

vars: {
  nick: string | *"dr_verm" @tag(nick)
  channel: string | *"#dr_verm" @tag(channel)
}

cfg: {
  today: "working on a twitch bot in CUEtils pipelines"
  music: "streaming https://www.youtube.com/watch?v=udGvUx70Q3U"
}

meta: {
  @pipeline(meta)
  secrets: {
    env: { 
      TWITCH_IRCBOT_KEY: _ @task(os.Getenv)
    } 
    key: env.TWITCH_IRCBOT_KEY
  }

  irc: {
    nick: string & vars.nick 
    pass: string & secrets.key
    host: string & "irc.chat.twitch.tv:6667"
    channel: string & vars.channel
    log_msgs: bool | *true
    persistent_msglogs: bool | *true
    // are these messages standardized
    pong: string & "tmi.twitch.tv"

    init_msgs: [
      "JOIN " + channel,
      "CAP REQ :twitch.tv/membership",
      "CAP REQ :twitch.tv/tags",
      "CAP REQ :twitch.tv/commands",
    ]

    handler: IRCHandler

  }

}

listen: {
  @pipeline()

  cfg: meta

  bot: cfg.irc & {
    @task(msg.IrcClient)
  }
}

IRCHandler: {
  // input
  msg: _

  // output (oneof), first taken anyhow
  error?: _
  resp?: _
  pipe?: _

  if msg.Command == "PRIVMSG" {
    cmd: msg.Params[1]

    [
      if cmd == "!today" { resp: cfg.today },
      if cmd == "!music" { resp: cfg.music },
      if cmd == "!docker" { pipe: IRCHandlers.docker },
      if strings.HasPrefix(cmd, "!so") { pipe: IRCHandlers.shoutout & { args: cmd } },

      { error: "unknown cmd: " + cmd },
    ][0] 
  }

  if msg.Command == "USEREVENT" {

  }
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
