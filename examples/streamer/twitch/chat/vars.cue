package chat

vars: {
  nick: string | *"dr_verm" @tag(nick)
  channel: string | *"#dr_verm" @tag(channel)
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

