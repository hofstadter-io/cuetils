// This pipeline listens to the channel chat
// and does things, hopefully
package chat

import "strings"

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
      // if pipeHandlers[cmd] != _|_ {
      //   pipe: pipeHandlers[cmd] & { args: parts }
      // }

      { error: "unknown cmd: " + cmd },
    ] 

    switch[0]
  }

  if msg.Command == "USEREVENT" {

  }
}
