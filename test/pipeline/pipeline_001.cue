import "encoding/json"

req: {
	host: "https://postman-echo.com"
	method: "GET"
	path: "/get"
	query: {
		cow: "moo"
	}
}

pick: {
	args: cow: string
}

tasks: [string]: {
	Out: _
	...
}

apicall: {
  @pipeline(apicall)
	r1: { #Req: req, Resp: _ } @task(api/call) @print("#Req",Resp)
	p1: { #X: r1.Resp, #P: pick } @task(st/pick) @print(Out)
}

readfile: {
  @pipeline(readfile)
	r: { #F: "../tree.json", Contents: string } @task(os/readfile)
  j: json.Unmarshal(r.Contents)
  p: { #X: j, #P: { tree: cow: _ } } @task(st/pick)

  final: { #O: p.Out.tree } @task(os/stdout)
}
