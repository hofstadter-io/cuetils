package tests

import (
	st "github.com/hofstadter-io/cuetils/structural"
)

count_tests: [
	{
		tree: {
			a: {
				foo: "bar"
				a: b: c: "d"
			}
			cow: "moo"
		}

		#count: (st.#Count & { #in: tree }).count
		#real: 9
		same: #real & #count
	},
	{
		tree: {
			a: int
			b: [1,2,4]
		}

		#count: (st.#Count & { #in: tree }).count
		#real: 6
		same: #real & #count
	},
]
