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

	cuest, err := NewCuest("mask")
	if err != nil {
		return nil, err
	}

	origs, err := LoadInputs([]string{orig}, cuest)
	if len(origs) == 0 {
		return nil, fmt.Errorf("original found")
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

	// only handling one orig for now, fill into val beforehand
	ov := cuest.ctx.CompileBytes(origs[0].Content, cue.Filename(origs[0].Filename))
	if ov.Err() != nil {
		return nil, ov.Err()
	}
	// update val with the orig value
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
