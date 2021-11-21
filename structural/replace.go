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
	return BinaryOpGlobs(code, globs, opts, ReplaceValue)
}

func ReplaceValue(repl, val cue.Value, opts *flags.RootPflagpole) (cue.Value, error) {
	r, _ := replaceValue(repl, val, opts)
	return r, nil
}

func replaceValue(repl, target cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	switch target.IncompleteKind() {
	case cue.StructKind:
		return replaceStruct(repl, target, opts)

	case cue.ListKind:
		return replaceList(repl, target, opts)

	default:
		// should already have the same label by now
		// but maybe not if target is basic and repl is not
		return repl, true
	}
}

func replaceStruct(repl, target cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	ctx := target.Context()

	result := newStruct(ctx)
	iter, _ := target.Fields(defaultWalkOptions...)

	cnt := 0
	for iter.Next() {
		cnt++
		s := iter.Selector()
		p := cue.MakePath(s)
		r := repl.LookupPath(p)
		// fmt.Println(cnt, iter.Value(), f, f.Exists())
		// check that field exists in from. Should we be checking f.Err()?
		if r.Exists() {
			v, ok := replaceValue(r, iter.Value(), opts)
			// fmt.Println("r:", r, ok, p)
			if ok {
				result = result.FillPath(p, v)
			}
		} else {
			// include if not in replace
			result = result.FillPath(p, iter.Value())
		}
	}

	return result, true
}

func replaceList(repl, target cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	ctx := target.Context()

	ri, _ := repl.List()
	ti, _ := target.List()

	result := []cue.Value{}
	for ri.Next() && ti.Next() {
		r, ok := replaceValue(ri.Value(), ti.Value(), opts)
		if ok {
			result = append(result, r)
		}
	}
	return ctx.NewList(result...), true
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
