import "encoding/json"

apicall: {
  @pipeline(apicall)
  In: string
	r: { f: In, contents: string } @task(os.ReadFile)
  j: json.Unmarshal(r.contents)
	r1: { #Req: j, Resp: _ } @task(api.Call)
  Resp: r1.Resp
}

main: {
  @pipeline(main)

  start: { text: "apicalling" } @task(os.Stdout)

  call: apicall & { In: "req.json" } @dummy(call1)
  final: { text: call.Resp } @task(os.Stdout,final1) @print(text)

  call2: apicall & { In: "req2.json" } @dummy(call2)
  final2: { text: call2.r1.Resp } @task(os.Stdout,final2) @print(text)

}
