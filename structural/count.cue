package structural

import (
	"list"

	"github.com/hofstadter-io/cuetils/recurse"
)

#countF: {
  #next: _
	#func: {
    #in: _
		#multi: {...} | [...]
    count: {
			if (#in & #multi) == _|_ { 1 }
			if (#in & {...}) != _|_ {
				list.Sum([for k,v in #in {(#next & {#in: v}).count}]) + len(#in)
			}
			if (#in & [...]) != _|_ {
				list.Sum([for k,v in #in {(#next & {#in: v}).count}])
			}
    }
  }
}

#Count: recurse.#RecurseN & {#funcFactory: #countF}
