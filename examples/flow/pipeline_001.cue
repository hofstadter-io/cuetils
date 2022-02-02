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
  @flow()
	r1: { "req": req, resp: _ } @task(api.Call) @print(req,resp)
	p1: { "val": r1.resp, "pick": pick } @task(st.Pick) @print(out)
}

readfile: {
  @flow()
	r: { filename: "cue/req.json", contents: string } @task(os.ReadFile)
  j: json.Unmarshal(r.contents)
  p: { val: j, pick: { query: _ } } @task(st.Pick)

  final: { text: p.out } @task(os.Stdout)
}
