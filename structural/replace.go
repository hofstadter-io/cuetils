package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/errors"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

const replacefmt = `
val: #Replace%s
val: #X: _
val: #R: _
replace: val.replace
`

func ReplaceGlobs(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	return ReplaceGlobsCue(code, globs, rflags)
}

func ReplaceGlobsGo(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	return BinaryOpGlobs(code, globs, rflags, ReplaceValue)
}

func ReplaceValue(repl, val cue.Value) (cue.Value, error) {
	r, _ := replaceValue(repl, val)
	return r, nil
}

func replaceValue(repl, val cue.Value) (cue.Value, bool) {
	ctx := repl.Context()

	// the lhs/rhs values should have the same label / type
	// are we sure they do at this point?
	// TODO, make tests

	switch val.IncompleteKind() {
	case cue.StructKind:

	case cue.ListKind:
		lpt, err := getListProcType(repl)
		if err != nil {
			ce := errors.Newf(repl.Pos(), "%v", err)
			ev := ctx.MakeError(ce)
			return ev, true
		}

		_ = lpt

	default:
		// should already have the same label by now
		return repl, true
	}
	return val, false
}

func ReplaceGlobsCue(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	cuest, err := NewCuest([]string{"replace"}, nil)
	if err != nil {
		return nil, err
	}

	operator, err := ReadArg(code, cuest.ctx, nil)
	if err != nil {
		return nil, err
	}

	inputs, err := LoadGlobs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	// construct reusable val with function
	maxiter := ""
	if mi := rflags.Maxiter; mi > 0 {
		maxiter = fmt.Sprintf(" & { #maxiter: %d }", mi)
	}
	content := fmt.Sprintf(replacefmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	// fill val with the orig value, so we only need to once before loop
	val = val.FillPath(cue.ParsePath("val.#R"), operator.Value)

	results := make([]GlobResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#X"), iv)

		v := result.LookupPath(cue.ParsePath("replace"))

		results = append(results, GlobResult{
			Filename: input.Filename,
			Value:    v,
		})

	}

	return results, nil
}
