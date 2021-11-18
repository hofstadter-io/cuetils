package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

const extendfmt = `
val: #Extend%s
val: #X: _
val: #E: _
extend: val.extend
`

func ExtendGlobs(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	return ExtendGlobsCue(code, globs, rflags)
}

func ExtendGlobsGo(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	return BinaryOpGlobs(code, globs, rflags, ExtendValue)
}

func ExtendValue(ext, val cue.Value) (cue.Value, error) {
	r, _ := extendValue(ext, val)
	return r, nil
}

func extendValue(ext, val cue.Value) (cue.Value, bool) {
	return ext, false
}

func ExtendGlobsCue(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	cuest, err := NewCuest([]string{"extend"}, nil)
	if err != nil {
		return nil, err
	}

	operator, err := ReadArg(code, rflags.Load, cuest.ctx, nil)
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
	val = val.FillPath(cue.ParsePath("val.#E"), operator.Value)

	results := make([]GlobResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#X"), iv)

		v := result.LookupPath(cue.ParsePath("extend"))

		results = append(results, GlobResult{
			Filename: input.Filename,
			Value:    v,
		})

	}

	return results, nil
}
