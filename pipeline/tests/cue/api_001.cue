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
  @pipeline(api-test)
  In: { 
    "req": req
    "p": pick
  }
	r: { #Req: In.req, Resp: _ } @task(api.Call)
	o: { text: r.Resp } @task(os.Stdout)
}
