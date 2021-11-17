package structural

import (
	"github.com/hofstadter-io/cuetils/recurse"
)

#extendF: {
	#next: _
	#func: {
		#X: _
		#E: _
		extend: {
			for i,x in #X {
				let e = #E[i]
				if e != _|_ {
					if (x & {...}) != _|_ {
						"\(i)": (#next & { #X: x, #E: e }).extend
					}
					// keep anything in x
					if (x & {...}) == _|_ {
						"\(i)": x
					}
				}
				// keep anything in x
				if e == _|_ {
					"\(i)": x
				}
			}

			// now look for anything in U that is not in X
			{
				for i,e in #E {
					if #X[i] == _|_ {
						"\(i)": e
					}
				}
			}
		}
	}
}

#Extend: recurse.#RecurseN & { #funcFactory: #extendF }
