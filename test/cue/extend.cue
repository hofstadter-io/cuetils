package tests

import (
	st "github.com/hofstadter-io/cuetils/structural"
)

extend_tests: [
	{
		x: {
			a: "a"
			e: {
				a: "a"
				b: "b"
			}
		}
		e: {
			a: "A"
			b: "b"
			e: {
				b: 2
				c: "c"
			}
			d: int
		}
		#extend: (st.#Extend & { #X: x, #E: e }).extend
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
		same: #real & #extend
	},
]



