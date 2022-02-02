import "encoding/json"

vars: {
  which: string @tag(which)
}

apicall: {
  @pipeline(apicall)
  In: string
	r: { filename: In, contents: string } @task(os.ReadFile)
  j: json.Unmarshal(r.contents)
	r1: { req: j, resp: _ } @task(api.Call)
  Resp: r1.resp
}

main: {
  @pipeline()

  start: { text: "apicalling" } @task(os.Stdout)

  call: apicall & { In: "req.json" } @dummy(call1)
  final: { text: call.r1.resp } @task(os.Stdout,final1)

  call2: apicall & { In: "req2.json" } @dummy(call2)
  final2: { text: call2.Resp } @task(os.Stdout,final2)

}
