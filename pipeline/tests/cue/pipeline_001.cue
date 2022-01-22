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
  @pipeline()
	r1: { #Req: req, Resp: _ } @task(api.Call) @print("#Req",Resp)
	p1: { #X: r1.Resp, #P: pick } @task(st.Pick) @print(Out)
}

readfile: {
  @pipeline()
	r: { f: "../tree.json", contents: string } @task(os.ReadFile)
  j: json.Unmarshal(r.contents)
  p: { #X: j, #P: { tree: cow: _ } } @task(st.Pick)

  final: { text: p.Out.tree } @task(os.Stdout)
}
