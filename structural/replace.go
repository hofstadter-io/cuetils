package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

const replacefmt = `
val: #Replace%s
val: #X: _
val: #R: _
replace: val.replace
`

func ReplaceGlobs(code string, globs []string, opts *flags.RootPflagpole) ([]GlobResult, error) {
	return ReplaceGlobsCue(code, globs, opts)
}

func ReplaceGlobsGo(code string, globs []string, opts *flags.RootPflagpole) ([]GlobResult, error) {
	return BinaryOpGlobs(code, globs, opts, ReplaceValue)
}

// TODO, make opts a pointer so we are only passing that during recursion
func ReplaceValue(repl, val cue.Value, opts *flags.RootPflagpole) (cue.Value, error) {
	r, _ := replaceValue(repl, val, opts)
	return r, nil
}

func replaceValue(repl, val cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	switch val.IncompleteKind() {
	case cue.StructKind:
		return replaceStruct(repl, val, opts)

	case cue.ListKind:
		return replaceList(repl, val, opts)

	default:
		// should already have the same label by now
		return repl, true
	}
	return val, false
}

func replaceStruct(repl, from cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {

	return repl, false
}

func replaceList(repl, from cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {

	return repl, false
}

func ReplaceGlobsCue(code string, globs []string, opts *flags.RootPflagpole) ([]GlobResult, error) {
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
	if mi := opts.Maxiter; mi > 0 {
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
