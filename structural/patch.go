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

func Patch(orig string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	cuest, err := NewCuest([]string{"patch"}, nil)
	if err != nil {
		return nil, err
	}

	operator, err := ParseOperator(orig)
	if err != nil {
		return nil, err
	}
	operator, err = LoadOperator(operator, rflags.LoadOperands, cuest.ctx)
	if err != nil {
		return nil, err
	}

	inputs, err := ReadGlobs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	// construct reusable val with function
	maxiter := ""
	if mi := rflags.Maxiter; mi > 0 {
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
