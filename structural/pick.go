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

func Pick(orig string, globs []string, rflags flags.RootPflagpole) ([]PickResult, error) {
	// no globs, then stdin
	if len(globs) == 0 {
		globs = []string{"-"}
	}

	cuest, err := NewCuest([]string{"pick"}, nil)
	if err != nil {
		return nil, err
	}

	ov, err := LoadInputs([]string{orig}, cuest.ctx)
	if err != nil {
		return nil, err
	}
	if ov.Err() != nil {
		return nil, ov.Err()
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
	val = val.FillPath(cue.ParsePath("val.#P"), ov)

	picks := make([]PickResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#X"), iv)

		dv := result.LookupPath(cue.ParsePath("pick"))

		out, err := FormatOutput(dv, rflags.Out)
		if err != nil {
			return nil, err
		}

		picks = append(picks, PickResult{
			Filename: input.Filename,
			Content:  out,
		})

	}

	return picks, nil
}
