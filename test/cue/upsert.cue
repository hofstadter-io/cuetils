package tests

import (
	st "github.com/hofstadter-io/cuetils/structural"
)

upsert_tests: [
	{
		x: {
			a: "a"
			b: "b"
			e: {
				a: "a"
				b: "b"
			}
		}
		u: {
			a: "A"
			e: {
				b: 2
				c: "c"
			}
			d: int
		}
		#upsert: (st.#Upsert & { #X: x, #U: u }).upsert
		#real: {
			a: "A"
			b: "b"
			e:  {
				a: "a"
				b: 2
				c: "c"
			}
			d: int
		}
		same: #real & #upsert
	},
]


