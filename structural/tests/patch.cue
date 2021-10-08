package tests

import (
	st "github.com/hofstadter-io/cuetils/structural"
)

patch_tests: [
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
		p: {
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
		patch: (st.#Patch & { #X: x, #P: p }).patch
		patch: {
			b: "b"
			c: "c"
			d: "D"
			e:  {
				b: "b"
				c: "c"
				d: 1
			}
		}
	}
]


