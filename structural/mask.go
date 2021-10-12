package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type MaskResult struct {
	Filename string
	Content  string
}

const maskfmt = `
val: #Mask%s
val: #X: _
val: #M: _
mask: val.mask
`

func Mask(orig string, globs []string) ([]MaskResult, error) {
	// no globs, then stdin
	if len(globs) == 0 {
		globs = []string{"-"}
	}

	cuest, err := NewCuest([]string{"mask"}, nil)
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
	content := fmt.Sprintf(maskfmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	// fill val with the orig value, so we only need to once before loop
	val = val.FillPath(cue.ParsePath("val.#M"), ov)

	masks := make([]MaskResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#X"), iv)

		dv := result.LookupPath(cue.ParsePath("mask"))

		out, err := FormatOutput(dv, flags.RootPflags.Out)
		if err != nil {
			return nil, err
		}

		masks = append(masks, MaskResult{
			Filename: input.Filename,
			Content:  out,
		})

	}

	return masks, nil
}
