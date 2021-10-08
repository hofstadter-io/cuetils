package structural

import (
	"github.com/hofstadter-io/cuetils/recurse"
)

#patchF: {
	#next: _
	#func: {
		#X: _
		#P: _
		patch: {
			{
				for i,x in #X {
					let pp = #P["+"][i]
					let pn = #P["-"][i]

					// in the "+" set
					if pp != _|_ {
						"\(i)": pp
					}
					// not in the "-" set, so conditionally add
					if pn == _|_ {
						let p = #P[i]
						if p != _|_ {
							// may not need this condition
							if (x & {...}) != _|_ {
								"\(i)": (#next & { #X: x, #P: p }).patch
							}
						}
						// if not found, it stayed the same and is a basic lit
						if p == _|_ {
							"\(i)": x			
						}
					}
					// else we aren't adding it

					// now add anything in '+' that isn't in x
					{
						for i,p in #P["+"] {
							if #X[i] == _|_ {
								"\(i)": p
							}
						}
					}
				}
			} 
		}
	}
}

#Patch: recurse.#RecurseN & { #funcFactory: #patchF }
