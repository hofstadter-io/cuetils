package structural

import (
	"list"

	"github.com/hofstadter-io/cuetils/recurse"
)

#depthF: {
  #next: _
	#func: {
    #in: _
		#basic: int|number|string|bytes|null
    out: {
			if (#in & #basic) != _|_ { 1 }
			if (#in & #basic) == _|_ {
				list.Max([for k,v in #in {(#next & {#in: v}).out}]) + 1
			}
    }
  }
}

#Depth: recurse.#RecurseN & {#funcFactory: #depthF}
