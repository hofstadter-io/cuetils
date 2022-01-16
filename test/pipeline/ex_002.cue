x: {
	a: {
		b: string
	}
	c: int
	d: "D"
}

y: {
	"a": {
		"b": "B"
	},
	"b": 1,
	"c": 2,
	"d": "D"
}

tasks: {
	p1: { #X: x, #P: y } @pick()
	u1: {
		@upsert()
		#X: { #X: p1, #M: { c: int } } @mask()
		#U: { #X: p1, #M: { a: _ } } @mask()
	}
}
