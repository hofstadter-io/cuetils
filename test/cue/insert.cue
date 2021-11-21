package tests

import (
	st "github.com/hofstadter-io/cuetils/structural"
)

insert_tests: [
	{
		x: {
			a: "a"
			e: {
				a: "a"
				b: "b"
			}
		}
		i: {
			a: "A"
			b: "b"
			e: {
				b: 2
				c: "c"
			}
			d: int
		}
		#insert: (st.#Insert & { #X: x, #I: i }).insert
		#real: {
			a: "a"
			b: "b"
			e:  {
				a: "a"
				b: "b"
				c: "c"
			}
			d: int
		}
		same: #real & #insert
	},
]



