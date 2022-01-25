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

tasks: {
  @pipeline()
  In: { 
    "req": req
    "p": pick
  }
	r: { req: In.req, resp: _ } @task(api.Call)
	o: { text: r.resp } @task(os.Stdout)
}
