@flow()

config: {
  port: "2323"
}

server: {
  @flow(server)

  wait: { duration: "60s", done: _ } @task(os.Sleep)

  run: {
    @task(api.Serve)
    w: wait.done

    port: config.port
    routes: {
      "/hello": {
        method: "GET"
        resp: {
          status: 200
          body: "hallo chat!"
        }
      }
    }
  }
}
