package structural

import (
	"github.com/hofstadter-io/cuetils/recurse"
)

#upsertF: {
	#next: _
	#func: {
		#X: _
		#U: _
		upsert: {
			for i,x in #X {
				let u = #U[i]
				if u != _|_ {
					if (x & {...}) != _|_ {
						"\(i)": (#next & { #X: x, #U: u }).upsert
					}
					if (x & {...}) == _|_ {
						"\(i)": u
					}
				}
				if u == _|_ {
					"\(i)": x
				}
			}

			// now look for anything in Y that is not in X
			{
				for i,u in #U {
					if #X[i] == _|_ {
						"\(i)": u
					}
				}
			}
		}
	}
}

#Upsert: recurse.#RecurseN & { #funcFactory: #upsertF }
