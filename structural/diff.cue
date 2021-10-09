package structural

import (
	"github.com/hofstadter-io/cuetils/recurse"
)

#diffF: {
	#next: _
	#func: {
		#X: _
		#Y: _
		diff: {
			{
				for i,x in #X {
					let y = #Y[i]
					// not in Y, so add to "-" set
					if y == _|_ {
						"-": "\(i)": x			
					}
					// in both
					if y != _|_ {
						// if struct, then recurse
						if (x & {...}) != _|_ {
							"\(i)": (#next & { #X: x, #Y: y }).diff
						}
						// not struct, so replace if doesn't unify
						if (x & {...}) == _|_ {
							// this is where we could add more complexity for `int & 1`
							if (x & y) == _|_ {
								"-": "\(i)": x
								"+": "\(i)": y
							}
						}
					}
				}
			}

			// now look for anything in Y that is not in X
			{
				for i,y in #Y {
					if #X[i] == _|_ {
						"+": "\(i)": y
					}
				}
			}
		}
	}
}

#Diff: recurse.#RecurseN & { #funcFactory: #diffF }
