package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type PickResult struct {
	Filename string
	Content  string
}

const pickfmt = `
val: #Pick%s
val: #X: _
val: #P: _
pick: val.pick
`

func Pick(pick string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	cuest, err := NewCuest([]string{"pick"}, nil)
	if err != nil {
		return nil, err
	}

	operator, err := ParseOperator(pick)
	if err != nil {
		return nil, err
	}
	operator, err = LoadOperator(operator, rflags.Load, cuest.ctx)
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
	content := fmt.Sprintf(pickfmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	// fill val with the orig value, so we only need to once before loop
	val = val.FillPath(cue.ParsePath("val.#P"), operator.Value)

	results := make([]GlobResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		// special case for pick
		if pick == "_" {
			results = append(results, GlobResult{
				Filename: input.Filename,
				Value:    iv,
			})
			continue
		}

		result := val.FillPath(cue.ParsePath("val.#X"), iv)

		v := result.LookupPath(cue.ParsePath("pick"))

		results = append(results, GlobResult{
			Filename: input.Filename,
			Value:    v,
		})

	}

	return results, nil
}
