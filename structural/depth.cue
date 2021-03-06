package structural

import (
	"list"

	"github.com/hofstadter-io/cuetils/recurse"
)

#depthF: {
  #next: _
	#func: {
    #in: _
		#multi: {...} | [...]
    depth: {
			if (#in & #multi) == _|_ { 1 }
			if (#in & {...}) != _|_ {
				list.Max([for k,v in #in {(#next & {#in: v}).depth}]) + 1
			}
			if (#in & [...]) != _|_ {
				list.Max([for k,v in #in {(#next & {#in: v}).depth}])
			}
    }
  }
}

#Depth: recurse.#RecurseN & {#funcFactory: #depthF}
