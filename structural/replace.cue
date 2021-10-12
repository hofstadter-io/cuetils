package structural

import (
	"github.com/hofstadter-io/cuetils/recurse"
)

#replaceF: {
	#next: _
	#func: {
		#X: _
		#R: _
		replace: {
			for i,x in #X {
				let r = #R[i]
				if r != _|_ {
					if (x & {...}) != _|_ {
						"\(i)": (#next & { #X: x, #R: r }).replace
					}
					if (x & {...}) == _|_ {
						"\(i)": r
					}
				}
				// todo, replace with the must operator
				if r == _|_ {
					"\(i)": x
				}
			}
		}
	}
}

#Replace: recurse.#RecurseN & { #funcFactory: #replaceF }
