package tests

import (
	st "github.com/hofstadter-io/cuetils/structural"
)

pick_tests: [
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
		p: {
			b: string
			d: int
			e: {
				a: _
				b: =~"^b"
				d: =~"^d"
			}
    }
		#pick: (st.#Pick & { #X: x, #P: p }).pick
		#real: {
			b: "b"
			e:  {
				a: "a"
				b: "b1"
			}
		}
		same: #real & #pick
	},
	{
		x: {
			c: {
				a: "a"
				b: "b"
				c: {
					a: "a"
					b: "b"
				}
			}
		}
		p: {
			c: {
				a: _
			}
    }
		#pick: (st.#Pick & { #X: x, #P: p }).pick
		#real: {
			c:  {
				a: "a"
			}
		}
		same: #real & #pick
	},
]

