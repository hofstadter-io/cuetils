import "encoding/json"

apicall: {
  @pipeline(apicall)
  In: string
	r: { #F: In, Contents: string } @task(os/readfile)
  j: json.Unmarshal(r.Contents)
	r1: { #Req: j, Resp: _ } @task(api/call) @print("#Req",Resp)
  Resp: r1.Resp
}


main: {
  @pipeline(main)

  input: { #O: "apicalling" } @task(os/stdout)

  call: apicall & { In: "req.json" }
  final: { #O: call.Resp } @task(os/stdout)

  call2: apicall & { In: "req2.json" }
  final2: { #O: call2.Resp } @task(os/stdout)
}

