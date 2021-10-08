package tests

import (
	st "github.com/hofstadter-io/cuetils/structural"
)

mask_tests: [
	{
		x: {
			a: "a"
			b: "b"
			d: "d"
			e: {
				a: "a"
				b: "b1"
				d: "cd"
			}
		}
		m: {
			b: string
			d: int
			e: {
				a: _
				b: =~"^b"
				d: =~"^d"
			}
		}
		#mask: (st.#Mask & { #X: x, #M: m }).mask
		#real: {
			a: "a"
			d: "d"
			e:  {
				d: "cd"
			}
		}
		same: #real & #mask
	},
]


