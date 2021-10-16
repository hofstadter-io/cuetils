package tests

import (
	st "github.com/hofstadter-io/cuetils/structural"
)

diff_tests: [
	{
		x: {
			a: "a"
			b: "b"
			d: "d"
			e: {
				a: "a"
				b: "b"
				d: "d"
			}
		}
		y: {
			b: "b"
			c: "c"
			d: "D"
			e:  {
				b: "b"
				c: "c"
				d: 1
			}
		}
		#diff: (st.#Diff & { #X: x, #Y: y }).diff
		#real: {
			"-": {
				a: "a"
				d: "d"
			}
			e: {
				"-": {
					a: "a"
					d: "d"
				}
				"+": {
					d: 1
					c: "c"
				}
			}
			"+": {
				d: "D"
				c: "c"
			}
    }
		same: #real & #diff
	}
]

