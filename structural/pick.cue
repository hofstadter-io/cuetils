package structural

import (
	"github.com/hofstadter-io/cuetils/recurse"
)

#pickF: {
	#next: _
	#func: {
		#X: _
		#P: _
		pick: {
			for i,p in #P {
				let x = #X[i]
				// in both
				if x != _|_ {
					// if they unify, then just add
					if (x & p) != _|_ {
						"\(i)": x
					}
					// if they do not unify
					if (x & p) == _|_ {
						// and if struct, then recurse
						if (x & {...}) != _|_ {
							"\(i)": (#next & { #X: x, #P: p }).pick
						}
					}
				}
			}
		}
	}
}

#Pick: recurse.#RecurseN & { #funcFactory: #pickF }
