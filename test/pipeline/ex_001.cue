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

tasks: {
	p1: { #X: x, #P: y } @task(st/pick)
}
