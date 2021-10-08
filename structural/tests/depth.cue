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

		depth: (st.#Depth & { #in: tree }).out
		real: 5
		same: real & depth
	},
]
