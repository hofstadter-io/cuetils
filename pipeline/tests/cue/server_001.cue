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
      resp: "hallo chat!"
    }
    "/echo": {
      method: "GET"
      req: _
      resp: req
      // resp: req.query.cow
    }
    "/pipe": {
      @pipeline()
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
