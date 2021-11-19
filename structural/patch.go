package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

const patchfmt = `
val: #Patch%s
val: #X: _
val: #P: _
patch: val.patch
`

func PatchGlobs(patch string, orig string, opts *flags.RootPflagpole) ([]GlobResult, error) {
	return PatchGlobsCue(patch, orig, opts)
}

func PatchGlobsGo(patch string, orig string, opts *flags.RootPflagpole) ([]GlobResult, error) {
	return BinaryOpGlobs(patch, []string{orig}, opts, PatchValue)
}

func PatchValue(patch, val cue.Value, opts *flags.RootPflagpole) (cue.Value, error) {
	r, _ := patchValue(patch, val, opts)
	return r, nil
}

func patchValue(patch, val cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	return val, false
}

func PatchGlobsCue(patch string, orig string, opts *flags.RootPflagpole) ([]GlobResult, error) {
	cuest, err := NewCuest([]string{"patch"}, nil)
	if err != nil {
		return nil, err
	}

	operator, err := ReadArg(patch, cuest.ctx, nil)
	if err != nil {
		return nil, err
	}

	inputs, err := LoadGlobs([]string{orig})
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	// construct reusable val with function
	maxiter := ""
	if mi := opts.Maxiter; mi > 0 {
		maxiter = fmt.Sprintf(" & { #maxiter: %d }", mi)
	}
	content := fmt.Sprintf(patchfmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	// update val with the orig value
	val = val.FillPath(cue.ParsePath("val.#P"), operator.Value)

	results := make([]GlobResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#X"), iv)

		v := result.LookupPath(cue.ParsePath("patch"))

		results = append(results, GlobResult{
			Filename: input.Filename,
			Value:    v,
		})

	}

	return results, nil
}
