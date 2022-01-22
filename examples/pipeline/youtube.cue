import "encoding/json"

@pipeline()

yt_api: "https://www.googleapis.com/youtube/v3"
channel: "UC9OMpomeCDawRQ9hAcLKMsw"

init: {
  env: { HOF_YOUTUBE_APIKEY: _ } @task(os.Getenv)
}

_secrets: {
  _key: init.env.HOF_YOUTUBE_APIKEY
}


info: {
  call: {
    @task(api.Call)
    _secrets
    _parts: "id,snippet"
    req: {
      host: yt_api
      method: "GET"
      _path: "/channels?id=\(channel)&part=\(_parts)"
      path: "\(_path)&key=\(call._key)"
    }
  }
  // print: { text: json.Indent(json.Marshal(call.resp), "", "  ") + "\n" } @task(os.Stdout)
}

playlists: {
  call: {
    @task(api.Call)
    _secrets
    req: {
      _parts: "id,snippet"
      host: yt_api
      method: "GET"
      path: "/playlists?channelId=\(channel)&part=\(_parts)&key=\(call._key)"
    }
  }
  pls: [ for item in call.resp.items { 
    id: item.id
    title: item.snippet.title
  }]
}

details: {
  _secrets
  for playlist in playlists.call.resp.items {
    _id: playlist.id
    call: (_id): { 
      @task(api.Call)
      req: {
        _parts: "id,snippet"
        host: yt_api
        method: "GET"
        _path: "/playlistItems?playlistId=\(_id)&part=\(_parts)"
        path: "\(_path)&key=\(details._key)"
      }
    }
    info: (_id): [ for item in details.call["\(_id)"].resp.items {
      id: item.id
      title: item.snippet.title
    }]
  }
}

final: {
  data: {
    username: info.call.resp.items[0].snippet.title
    for pl in playlists.pls {
      (pl.title): pl & { videos: details.info["\(pl.id)"] }
    }
  }
  print: { text: json.Indent(json.Marshal(data), "", "  ") + "\n" } @task(os.Stdout)
}
