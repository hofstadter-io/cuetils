// This pipeline listens to the channel chat
// and does things, hopefully
package twitch

import "strings"

vars: {
  nick: string | *"dr_verm" @tag(nick)
  channel: string | *"#dr_verm" @tag(channel)
}

links: {
  github:   "https://github.com/hofstadter-io"
  hof:      "https://github.com/hofstadter-io/hof"
  cuetils:  "https://github.com/hofstadter-io/cuetils"
  neoverm:  "https://github.com/verdverm/neoverm"
  grasshopper: "https://grasshopper.app/"
}

respHandlers: {
  "!today":    "working on a twitch bot in CUEtils pipelines"
  "!music":    "streaming https://www.youtube.com/watch?v=udGvUx70Q3U"
  "!github":   links.github 
  "!hof":      links.hof 
  "!cuetils":  links.cuetils 
  "!vim":      links.neoverm
  "!nvim":     links.neoverm
  "!neovim":   links.neoverm
  "!neoverm":  links.neoverm
  "!dox":      "google 'verdverm'"

  "!grasshopper": links.grasshopper 
  "!learn/code": "Try out the Grasshopper App \(links.grasshopper)"
  "!learn/go": "Checkout my currated links: https://verdverm.com/resources/learning-go"
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
    msg_cmd: msg.Params[1]
    parts: strings.Split(msg_cmd, " ")
    cmd: parts[0]

    switch: [
      if respHandlers[cmd] != _|_ {
        resp: respHandlers[cmd]
      }
      if pipeHandlers[cmd] != _|_ {
        pipe: pipeHandlers[cmd] & { args: parts }
      }

      { error: "unknown cmd: " + cmd },
    ] 

    switch[0]
  }

  if msg.Command == "USEREVENT" {

  }
}

pipeHandlers: {
  "!docker": {
    @pipeline()

    args?: [...string]

    get: {
      @task(os.Exec)
      cmd: ["docker", "ps"]
      stdout: string  
    }

    resp: strings.Replace(get.stdout, "\n", "", -1)
  }

  "!k8s": {
    @pipeline()

    args?: [...string]

    get: {
      @task(os.Exec)
      cmd: ["kubectl", "get", "pods", "--all-namespaces"]
      stdout: string  
    }

    // chill: { duration: "4s" } @task(os.Sleep)

    lines: strings.Split(get.stdout, "\n")
    count: len(lines) - 1

    resp: "there are \(count) pods running in Dr Verm's cluster" 
  }

  "!so": {
    // @pipeline()

    args?: [...string]
    who: args[1] 

    get: {
      // call twitch api for info about the user
      // eventually also look up custom data in DB
    }
    // chill: { duration: "4s" } @task(os.Sleep)

    resp: "you're the best \(who)"

  }
}
