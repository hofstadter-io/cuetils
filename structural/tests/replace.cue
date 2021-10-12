package tests

import (
	st "github.com/hofstadter-io/cuetils/structural"
)

replace_tests: [
	{
		x: {
			a: "A"
		}
		r: {
			a: int
		}
		#replace: (st.#Replace & { #X: x, #R: r }).replace
		#real: {
			a: int
		}
		// same: #real & #replace
	}, {
		x: {
			a: "a"
			b: "b"
			e: {
				a: "a"
				b: "b"
			}
		}
		r: {
			a: "A"
			e: {
				b: 2
				c: "c"
			}
		}
		#replace: (st.#Replace & { #X: x, #R: r }).replace
		#real: {
			a: "A"
			b: "b"
			e:  {
				a: "a"
				b: 2
			}
		}
		same: #real & #replace
	},
]


