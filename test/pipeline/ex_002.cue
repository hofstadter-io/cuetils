x: {
	"a": {
		"b": "B"
	},
	"b": 1
	"c": 2
	"d": "D"
}

y: {
	a: {
		b: string
	}
	c: int
	d: "D"
}

tasks: [string]: {
	Out: _
	...
}

tasks: {
	// @pipeline(main)

	p1: { #X: x, #P: y } @task(st/pick) @print(Out)
	
	m1: { #X: p1.Out, #M: { c: int } } @task(st/mask) @print(Out)
	m2: { #X: p1.Out, #M: { a: _ } } @task(st/mask) @print(Out)

	u1: { #X: m1.Out, #U: m2.Out } @task(st/upsert) @print(Out)
}
