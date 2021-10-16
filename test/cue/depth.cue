package tests

import (
	st "github.com/hofstadter-io/cuetils/structural"
)

depth_tests: [
	{
		tree: {
			a: {
				foo: "bar"
				a: b: c: "d"
			}
			cow: "moo"
		}

		#depth: (st.#Depth & { #in: tree }).depth
		#real: 5
		same: #real & #depth
	},
	{
		tree: {
			a: int
			b: [1,2,4]
			c: [{ a: b: "c" }]
		}

		#depth: (st.#Depth & { #in: tree }).depth
		#real: 4
		same: #real & #depth
	},
]
