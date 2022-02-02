import "encoding/json"

vars: {
  which: string @tag(which)
}

secrets: {
  user: "foobar" @secret()
}

apicall: {
  @flow(apicall)
  In: string
	r: { filename: In, contents: string } @task(os.ReadFile)
  j: json.Unmarshal(r.contents)
	r1: { req: j, resp: _ } @task(api.Call)
  Resp: r1.resp
}

main: {
  @flow()

  start: { text: "apicalling" } @task(os.Stdout)

  call: apicall & { 
    In: "req.json"
    key: "shhhh" @secret()
  }
  final: { text: call.r1.resp } @task(os.Stdout,final1)

}
