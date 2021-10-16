package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

const difffmt = `
val: #Diff%s
val: #X: _
val: #Y: _
diff: val.diff
`

func Diff(diff string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	cuest, err := NewCuest([]string{"diff"}, nil)
	if err != nil {
		return nil, err
	}

	operator, err := ParseOperator(diff)
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
	content := fmt.Sprintf(difffmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	// update val with the orig value
	val = val.FillPath(cue.ParsePath("val.#X"), operator.Value)

	results := make([]GlobResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#Y"), iv)

		v := result.LookupPath(cue.ParsePath("diff"))

		results = append(results, GlobResult{
			Filename: input.Filename,
			Value:  v,
		})
	}

	return results, nil
}
