import "encoding/json"

apicall: {
  @pipeline(apicall)
  In: string
	r: { filename: In, contents: string } @task(os.ReadFile)
  j: json.Unmarshal(r.contents)
	r1: { req: j, resp: _ } @task(api.Call)
  Out: r1.resp
}

main: {
  @pipeline(main)

  start: { text: "apicalling\n" } @task(os.Stdout)

  call: apicall & { In: "req.json" } @dummy(call1)
  final: { text: "final: " + json.Indent(json.Marshal(call.Out) + "\n", "", "  ") } @task(os.Stdout,final1)

  call2: apicall & { In: "req2.json" } @dummy(call2)
  final2: { text: "final2: " + json.Indent(json.Marshal(call2.Out) + "\n", "", "  ") } @task(os.Stdout,final2)
}
