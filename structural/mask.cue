package structural

import (
	"github.com/hofstadter-io/cuetils/recurse"
)

#maskF: {
	#next: _
	#func: {
		#X: _
		#M: _
		mask: {
			for i,x in #X {
				let m = #M[i]
				if (m&x) == _|_ {
					if (x & {...}) == _|_ {
						"\(i)": x
					}
					if (x & {...}) != _|_ {
						"\(i)": (#next & { #X: x, #M: m }).mask
					}
				}
				if (m&x) != _|_ {
					if (x & {...}) != _|_ {
						"\(i)": (#next & { #X: x, #M: m }).mask
					}
				}
			}
		}
	}
}

#Mask: recurse.#RecurseN & { #funcFactory: #maskF }

