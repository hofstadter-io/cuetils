import "encoding/json"

// This pipeline fetches a YouTube channel's
// [about, playlists, and videos] and
// combines them into a singular, brief representation

@pipeline()

yt_api: "https://www.googleapis.com/youtube/v3"
// channel: string @tag(channel)
channel: "UC9OMpomeCDawRQ9hAcLKMsw" @tag(channel)


// try to combine this with secrets
init: {
  env: { HOF_YOUTUBE_APIKEY: _ } @task(os.Getenv)
}
secrets: {
  key: init.env.HOF_YOUTUBE_APIKEY
}

info: {
  call: {
    @task(api.Call)
    parts: "id,snippet"
    req: {
      host: yt_api
      method: "GET"
      path: "/channels?id=\(channel)&part=\(parts)&key=\(secrets.key)"
    }
  }
  // print: { text: json.Indent(json.Marshal(call.resp), "", "  ") + "\n" } @task(os.Stdout)
}

playlists: {
  call: {
    @task(api.Call)
    parts: "id,snippet"
    req: {
      host: yt_api
      method: "GET"
      path: "/playlists?channelId=\(channel)&part=\(parts)&key=\(secrets.key)"
      // path: "\(_path)"
    }
  }
  // print: { text: json.Indent(json.Marshal(call.resp), "", "  ") + "\n" } @task(os.Stdout)
  pls: [ for item in call.resp.items { 
    id: item.id
    title: item.snippet.title
    link: "https://youtube.com/platlist?list=\(item.id)"
  }]
}

details: {
  for playlist in playlists.call.resp.items {
    _id: playlist.id
    call: (_id): { 
      @task(api.Call)
      parts: "id,snippet"
      req: {
        host: yt_api
        method: "GET"
        _path: "/playlistItems?playlistId=\(_id)&part=\(parts)"
        path: "\(_path)&key=\(secrets.key)"
      }
    }
    // print: (_id): { text: json.Indent(json.Marshal(call["\(_id)"].resp), "", "  ") + "\n" } @task(os.Stdout)
    info: (_id): [ for item in details.call["\(_id)"].resp.items {
      id: item.id
      title: item.snippet.title
      link: "https://youtu.be/\(item.snippet.resourceId.videoId)"
    }]
  }
}

final: {
  data: {
    username: info.call.resp.items[0].snippet.title
    "channel": "https://youtube.com/channel/\(channel)"
    for pl in playlists.pls {
      (pl.title): pl & { videos: details.info["\(pl.id)"] }
    }
  }
  print: { text: json.Indent(json.Marshal(data), "", "  ") + "\n" } @task(os.Stdout)
}
