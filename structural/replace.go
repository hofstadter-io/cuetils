package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type ReplaceResult struct {
	Filename string
	Content  string
}

const replacefmt = `
val: #Replace%s
val: #X: _
val: #R: _
replace: val.replace
`

func Replace(orig string, globs []string) ([]ReplaceResult, error) {
	// no globs, then stdin
	if len(globs) == 0 {
		globs = []string{"-"}
	}

	cuest, err := NewCuest([]string{"replace"}, nil)
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
	if mi := flags.RootPflags.Maxiter; mi > 0 {
		maxiter = fmt.Sprintf(" & { #maxiter: %d }", mi)
	}
	content := fmt.Sprintf(replacefmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	// fill val with the orig value, so we only need to once before loop
	val = val.FillPath(cue.ParsePath("val.#R"), ov)

	replaces := make([]ReplaceResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#X"), iv)

		dv := result.LookupPath(cue.ParsePath("replace"))

		out, err := FormatOutput(dv, flags.RootPflags.Out)
		if err != nil {
			return nil, err
		}

		replaces = append(replaces, ReplaceResult{
			Filename: input.Filename,
			Content:  out,
		})

	}

	return replaces, nil
}
