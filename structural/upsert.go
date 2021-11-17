package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

const upsertfmt = `
val: #Upsert%s
val: #X: _
val: #U: _
upsert: val.upsert
`

func UpsertGlobs(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	return UpsertGlobsCue(code, globs, rflags)
}

func UpsertGlobsGo(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	return nil, nil
}

func UpsertGlobsCue(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	cuest, err := NewCuest([]string{"upsert"}, nil)
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
	content := fmt.Sprintf(upsertfmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	// fill val with the orig value, so we only need to once before loop
	val = val.FillPath(cue.ParsePath("val.#U"), operator.Value)

	results := make([]GlobResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#X"), iv)

		v := result.LookupPath(cue.ParsePath("upsert"))

		results = append(results, GlobResult{
			Filename: input.Filename,
			Value:    v,
		})

	}

	return results, nil
}

func UpsertValue(up, val cue.Value) cue.Value {
	r, _ := upsertValue(up, val)
	return r
}

func upsertValue(up, val cue.Value) (cue.Value, bool) {
	return val, false
}
