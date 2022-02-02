import "encoding/json"

@pipeline()

config: {
  port: "2323"
}

run: {
  @task(api.Serve)

  port: config.port
  routes: {
    "/hello": {
      method: "GET"
      resp: {
        status: 200
        body: "hallo chat!"
      }
    }
    "/echo": {
      method: ["get", "post"]
      req: _
      // resp: req.query
      resp: json: req.query.cow
    }
    "/pipe": {
      @pipeline()
      req: _
      r: { filename: req.query.filename[0], contents: string } @task(os.ReadFile)
      j: json.Unmarshal(r.contents)
      resp: {
        status: 200
        json: j
      }
    }
  }
}

msg: {
  @task(os.Stdout)
  text: "should not get here either"
  dep: run.done
}

call: {
  wait: {
    @task(os.Sleep)
    duration: "2s"
    done: _
  }
  do: {
    @task(api.Call)
    dep: wait.done
    resp: string
    req: {
      host: "http://localhost:\(config.port)"
      method: "GET"
      path: "/hello"
      query: {
        cow: "moo"
      }
    }
  }
  say: {
    @task(os.Stdout)
    text: do.resp
  }
}

stop: {
  wait: {
    // @task(os.Sleep)
    duration: "2m"
    done: _
  }

  send: {
    // @task(msg.Chan)
    // to:
    // msg: "stop"
  }

}
