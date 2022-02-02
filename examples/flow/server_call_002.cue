@pipeline()

config: {
  port: "2323"
}

call: {
  do: {
    @task(api.Call)
    resp: string
    req: {
      host: "http://localhost:\(config.port)"
      method: "GET"
      path: "/hello"
      retry: {
        codes: [200,500]
        count: 5
      }
    }
  }
  say: {
    @task(os.Stdout)
    text: do.resp + "\n"
  }
}

