package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
)

const val = `
a: 1
b: "b"
c: "c" @error(my custom error message)
d: "d" @error()
`

func main() {
	ctx := cuecontext.New()

	v := ctx.CompileString(val, cue.Filename("val.cue"))
	fmt.Printf("%# v\n\n", v)

	b := v.LookupPath(cue.ParsePath("b"))
	v = v.FillPath(cue.ParsePath("b"), makeErr(b))

	c := v.LookupPath(cue.ParsePath("c"))
	v = v.FillPath(cue.ParsePath("c"), makeErr(c))

	d := v.LookupPath(cue.ParsePath("d"))
	v = v.FillPath(cue.ParsePath("d"), makeErr(d))

	err := v.Validate()
	fmt.Println(errors.Details(err, nil))
}

func makeErr(val cue.Value) cue.Value {
	var err error
	// create a message, inspect for @error(<msg>)
	msg := "my default error message"
	attr := val.Attribute("error")
	if attr.Err() == nil {
		if attr.NumArgs() == 0 {
			msg = "@error() requires contents for custom message"
		} else {
			msg, err = attr.String(0)
			if err != nil || msg == "" {
				msg = "@error() requires contents for custom message"
			}
		}
	}

	// create an error
	ce := errors.Newf(val.Pos(), msg)
	ev := val.Context().MakeError(ce)
	return ev
}
