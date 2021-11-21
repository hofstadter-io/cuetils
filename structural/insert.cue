package structural

import (
	"github.com/hofstadter-io/cuetils/recurse"
)

#insertF: {
	#next: _
	#func: {
		#X: _
		#I: _
		insert: {
			for l,x in #X {
				let i = #I[l]
				if i != _|_ {
					if (x & {...}) != _|_ {
						"\(l)": (#next & { #X: x, #I: i }).insert
					}
					// keep anything in x
					if (x & {...}) == _|_ {
						"\(l)": x
					}
				}
				// keep anything in x
				if i == _|_ {
					"\(l)": x
				}
			}

			// now look for anything in E that is not in X
			{
				for l,i in #I {
					if #X[l] == _|_ {
						"\(l)": i
					}
				}
			}
		}
	}
}

#Insert: recurse.#RecurseN & { #funcFactory: #insertF }
