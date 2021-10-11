# Cuetils

CUE utilities and helpers
for working with tree based objects
in any combination of CUE, Yaml, and JSON.

## Using

### As a binary

The `cuetils` program is useful for bulk operations
or when you don't want to write extra CUE or Go.

Download a [release from GitHub](https://github.com/hofstadter-io/cuetils/releases).

```
cuetils -h
```

### As a Go library

The Go libraries are optimized and have more capabilities.

```shell
go get github.com/hofstadter-io/cuetils@latest
```

```
import "github.com/hofstadter-io/cuetils/structural"
```

### As a CUE library

The CUE libraries all use the [RecurseN](#recursen) helper
and make use of the ["function" pattern](https://cuetorials.com/patterns/functions/).
You can also write [custom operators](#custom-operators).

Add to your project with `hof mod` or another method.
See: https://cuetorials.com/first-steps/modules-and-packages/#dependency-management

```
import "github.com/hofstadter-io/cuetils/structural"
```



## Structural Helpers


- [Count](#count) the nodes in an object
- [Depth](#depth) how deep an object is
- [Diff](#diff) two objects, producing a structured diff
- [Patch](#patch) an object, producing a new object
- [Pick](#pick) a subojbect from another, selecting only the parts you want
- [Maks](#mask) a subobject from another, filtering out parts you don't want
- [Replace](#replace) with a subobject, updating fields found
- [Upsert](#upsert) with a subobject, updating and adding fields
- [Transform](#transform) one or more objects into another using CUE
- [Validate](#validate) one or more objects with the power of CUE


Some notes:

The helpers work by checking if two operands unify.
We try to make note of the edge cases, which also
depends on if you are using CUE, Go, or `cuetils`.

- for __diff__ and __patch__, `int & 1` like expressions will _not_ be detected
- for __pick__ and __mask__, `int & 1` like expression are detected

Lists are not currently supported for __diff__ and __patch__.
It may be workable if the list sizes are known and order consistent.
[Associative Lists](https://cuetorials.com/cueology/futurology/associative-lists/)
may solve this issue. We don't currently have good syntax for specifying the key to match elements on.


### Count

__#Count__ calculates how many nodes are in an object.

<details>
<summary>Cuetils example</summary>
<br>

```cue
a: {
	foo: "bar"
	a: b: c: "d"
}
cow: "moo"
```

```shell
$ cuetils count tree.cue
9
```
</details>

<details>
<summary>CUE example</summary>
<br>

```cue
import "github.com/hofstadter-io/cuetils/structural"

tree: {
	a: {
		foo: "bar"
		a: b: c: "d"
	}
	cow: "moo"
}

depth: (structural.#Count & { #in: tree }).out
depth: 9
```
</details>

<details>
<summary>Go example</summary>
<br>
</details>


### Depth

__#Depth__ calculates the deepest branch of an object.

<details>
<summary>Cuetils example</summary>
<br>

```cue
a: {
	foo: "bar"
	a: b: c: "d"
}
cow: "moo"
```

```shell
$ cuetils depth tree.cue
5
```
</details>

<details>
<summary>CUE example</summary>
<br>

```cue
import "github.com/hofstadter-io/cuetils/structural"

tree: {
	a: {
		foo: "bar"
		a: b: c: "d"
	}
	cow: "moo"
}

depth: (structural.#Depth & { #in: tree }).out
depth: 5
```
</details>

<details>
<summary>Go example</summary>
<br>

```Go
import "github.com/hofstadter-io/cuetils/structural"
```
</details>


### Diff

__#Diff__ computes a semantic diff object

<details>
<summary>Cuetils example</summary>
<br>

```shell

```
</details>

<details>
<summary>CUE example</summary>
<br>

```cue
import "github.com/hofstadter-io/cuetils/structural"

x: {
	a: "a"
	b: "b"
	d: "d"
	e: {
		a: "a"
		b: "b"
		d: "d"
	}
}

y: {
	b: "b"
	c: "c"
	d: "D"
	e:  {
		b: "b"
		c: "c"
		d: 1
	}
}

diff: (structural.#Diff & { #X: x, #Y: y }).diff
diff: {
	"-": {
		a: "a"
		d: "d"
	}
	e: {
		"-": {
			a: "a"
			d: "d"
		}
		"+": {
			d: 1
			c: "c"
		}
	}
	"+": {
		d: "D"
		c: "c"
	}
}
```
</details>

<details>
<summary>Go example</summary>
<br>

```Go
import "github.com/hofstadter-io/cuetils/structural"
```
</details>


### Patch

__#Patch__ applies a diff object

<details>
<summary>Cuetils example</summary>
<br>

```shell

```
</details>

<details>
<summary>CUE example</summary>
<br>

```cue
import "github.com/hofstadter-io/cuetils/structural"

x: {
	a: "a"
	b: "b"
	d: "d"
	e: {
		a: "a"
		b: "b"
		d: "d"
	}
}

p: {
	"-": {
		a: "a"
		d: "d"
	}
	e: {
		"-": {
			a: "a"
			d: "d"
		}
		"+": {
			d: 1
			c: "c"
		}
	}
	"+": {
		d: "D"
		c: "c"
	}
}

patch: (structural.#Patch & { #X: x, #Y: y }).patch
patch: {
	b: "b"
	c: "c"
	d: "D"
	e:  {
		b: "b"
		c: "c"
		d: 1
	}
}
```
</details>

<details>
<summary>Go example</summary>
<br>

```Go
import "github.com/hofstadter-io/cuetils/structural"
```
</details>

### Pick

__#Pick__ extracts a subobject

<details>
<summary>Cuetils example</summary>
<br>

```shell

```
</details>

<details>
<summary>CUE example</summary>
<br>

```cue
import "github.com/hofstadter-io/cuetils/structural"

x: {
	a: "a"
	b: "b"
	d: "d"
	e: {
		a: "a"
		b: "b1"
		d: "cd"
	}
}
p: {
	b: string
	d: int
	e: {
		a: _
		b: =~"^b"
		d: =~"^d"
	}
}
pick: (structural.#Pick & { #X: x, #P: p }).pick
pick: {
	b: "b"
	e:  {
		a: "a"
		b: "b1"
	}
}
```
</details>

<details>
<summary>Go example</summary>
<br>

```Go
import "github.com/hofstadter-io/cuetils/structural"
```
</details>


### Mask

__#Mask__ removes a subobject

<details>
<summary>Cuetils example</summary>
<br>

```shell

```
</details>

<details>
<summary>CUE example</summary>
<br>

```cue
import "github.com/hofstadter-io/cuetils/structural"

x: {
	a: "a"
	b: "b"
	d: "d"
	e: {
		a: "a"
		b: "b1"
		d: "cd"
	}
}
m: {
	b: string
	d: int
	e: {
		a: _
		b: =~"^b"
		d: =~"^d"
	}
}
mask: (structural.#Mask & { #X: x, #M: m }).mask
mask: {
	a: "a"
	d: "d"
	e:  {
		d: "cd"
	}
}
```
</details>

<details>
<summary>Go example</summary>
<br>

```Go
import "github.com/hofstadter-io/cuetils/structural"
```
</details>


### Replace

<details>
<summary>Cuetils example</summary>
<br>

```shell

```
</details>

<details>
<summary>CUE example</summary>
<br>

```cue
import "github.com/hofstadter-io/cuetils/structural"
```
</details>

<details>
<summary>Go example</summary>
<br>

```Go
import "github.com/hofstadter-io/cuetils/structural"
```
</details>

### Upsert


<details>
<summary>Cuetils example</summary>
<br>

```shell

```
</details>

<details>
<summary>CUE example</summary>
<br>

```cue
import "github.com/hofstadter-io/cuetils/structural"
```
</details>

<details>
<summary>Go example</summary>
<br>

```Go
import "github.com/hofstadter-io/cuetils/structural"
```
</details>


### Transform


<details>
<summary>Cuetils example</summary>
<br>

```shell

```
</details>

<details>
<summary>Go example</summary>
<br>

```Go
import "github.com/hofstadter-io/cuetils/structural"
```
</details>


### Validate

<details>
<summary>Cuetils example</summary>
<br>

```shell

```
</details>

<details>
<summary>Go example</summary>
<br>

```Go
import "github.com/hofstadter-io/cuetils/structural"
```
</details>



### RecurseN

A function factory for bounded recursion, defaulting to 20.
This is a pattern to get around CUE's cycle detection
by creating a struct with fields named for each iteration.
See https://cuetorials.com/deep-dives/recursion/ for more details.

```cue
#RecurseN: {
	#maxiter: uint | *20
	#funcFactory: {
		#next: _
		#func: _
	}

	for k, v in list.Range(0, #maxiter, 1) {
		#funcs: "\(k)": (#funcFactory & {#next: #funcs["\(k+1)"]}).#func
	}
	#funcs: "\(#maxiter)": null

	#funcs["0"]
}
```

The core of the bounded recursion is this structural comprehension (unrolled for loop).
The "recursive call" is made with the following pattern.

```cue
(#funcFactory & {#next: #funcs["\(k+1)"]}).#func
```

You can override the iterations with `{ #maxdepth: 100 }` at the point of
usage or by creating new helpers from the existing ones.
You may need to adjust this

- up for deep objects
- down if runtime is an issue

```cue
import "github.com/hofstadter-io/cuetils/structural"

#LeaguesDeep: structural.#Depth & { #maxdepth: 10000 }
```


### Custom Helpers

You can make new helpers by building on the `#RecurseN` pattern.
You need two definitions, a factory and the user facing, recursed version.

```cue
package structural

import (
	"list"

	"github.com/hofstadter-io/cuetils/recurse"
)

// A function factory
#depthF: {
	// always required
	#next: _
	
	// the actual computation, must be named #func
	#func: {
		// you can have any args
		#in: _
		// or internal helpers
		#multi: {...} | [...]
		
		// the result, can be named anything
    depth: {
			// detect leafs
			if (#in & #multi) == _|_ { 1 }
			// detect struct
			if (#in & {...}) != _|_ {
				list.Max([for k,v in #in {(#next & {#in: v}).depth}]) + 1
			}
			// detect list
			if (#in & [...]) != _|_ {
				list.Max([for k,v in #in {(#next & {#in: v}).depth}])
			}
    }
	}
}

// The user facing, recursed version
#Depth: recurse.#RecurseN & {#funcFactory: #depthF}
```

The core of the recursive calling is:

```cue
(#next & {#in: v}).depth
```

