package pipeline

import (
	"fmt"
	"sync"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/structural"
)

var lock sync.Mutex

type PickTask struct {
	X   cue.Value `cue: "#X"`
	P   cue.Value `cue: "#P"`
	Ret cue.Value
}

// Tasks must implement a Run func, this is where we execute our task
func (P *PickTask) Run(t *flow.Task, err error) error {
	lock.Lock()
	defer func() {
		lock.Unlock()
	}()

	fmt.Println("PT start")
	if err != nil {
		fmt.Println("Dep error", err)
	}

	// not sure this is OK, but the value which was used for this task
	val := t.Value()
	s, err := formatCue(val)
	if err != nil {
		fmt.Println("Fmt error", err)
	}
	fmt.Printf("PickTask: %v\n", s)

	x := val.LookupPath(cue.ParsePath("#X"))
	p := val.LookupPath(cue.ParsePath("#P"))

	fmt.Printf("x: %v\n", x)
	fmt.Printf("p: %v\n", p)

	r, err := structural.PickValue(p, x, nil)
	if err != nil {
		return err
	}

	fmt.Printf("r: %v\n", r)

	// Use fill to "return" a result to the workflow engine
	res := val.FillPath(cue.ParsePath("Out"), r)

	t.Fill(res)

	return err
}

type MaskTask struct {
	X   cue.Value `cue: "#X"`
	M   cue.Value `cue: "#M"`
	Ret cue.Value
}

// Tasks must implement a Run func, this is where we execute our task
func (M *MaskTask) Run(t *flow.Task, err error) error {
	lock.Lock()
	defer func() {
		lock.Unlock()
	}()

	fmt.Println("MT start")
	if err != nil {
		fmt.Println("Dep error", err)
	}

	// not sure this is OK, but the value which was used for this task
	val := t.Value()

	s, err := formatCue(val)
	if err != nil {
		fmt.Println("Fmt error", err)
	}
	fmt.Printf("MaskTask: %v\n", s)

	x := val.LookupPath(cue.ParsePath("#X"))
	m := val.LookupPath(cue.ParsePath("#M"))

	fmt.Printf("x: %v\n", x)
	fmt.Printf("m: %v\n", m)

	// Use fill to "return" a result to the workflow engine
	// t.Fill(next)

	r, err := structural.MaskValue(m, x, nil)
	if err != nil {
		return err
	}

	fmt.Printf("r: %v\n", r)

	// Use fill to "return" a result to the workflow engine
	res := val.FillPath(cue.ParsePath("Out"), r)

	t.Fill(res)

	return err
}

type UpsertTask struct {
	X   cue.Value `cue: "#X"`
	U   cue.Value `cue: "#U"`
	Ret cue.Value
}

// Tasks must implement a Run func, this is where we execute our task
func (U *UpsertTask) Run(t *flow.Task, err error) error {
	lock.Lock()
	defer func() {
		lock.Unlock()
	}()

	fmt.Println("UT start")
	if err != nil {
		fmt.Println("Dep error", err)
	}

	// not sure this is OK, but the value which was used for this task
	val := t.Value()
	s, err := formatCue(val)
	if err != nil {
		fmt.Println("Fmt error", err)
	}
	fmt.Printf("UpsertTask: %v\n", s)

	x := val.LookupPath(cue.ParsePath("#X"))
	u := val.LookupPath(cue.ParsePath("#U"))

	fmt.Printf("x: %v\n", x)
	fmt.Printf("u: %v\n", u)

	r, err := structural.UpsertValue(u, x, nil)
	if err != nil {
		return err
	}

	fmt.Printf("r: %v\n", r)

	// Use fill to "return" a result to the workflow engine
	res := val.FillPath(cue.ParsePath("Out"), r)

	t.Fill(res)

	return err
}
