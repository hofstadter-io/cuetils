package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type UpsertResult struct {
	Filename string
	Content  string
}

const upsertfmt = `
val: #Upsert%s
val: #X: _
val: #U: _
upsert: val.upsert
`

func Upsert(orig string, globs []string) ([]UpsertResult, error) {
	// no globs, then stdin
	if len(globs) == 0 {
		globs = []string{"-"}
	}

	origs, err := LoadInputs([]string{orig})
	if len(origs) == 0 {
		return nil, fmt.Errorf("original found")
	}

	inputs, err := LoadInputs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	cuest, err := NewCuest("upsert")
	if err != nil {
		return nil, err
	}

	// construct reusable val with function
	maxiter := ""
	if mi := flags.RootPflags.Maxiter; mi > 0 {
		maxiter = fmt.Sprintf(" & { #maxiter: %d }", mi)
	}
	content := fmt.Sprintf(upsertfmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	// only handling one orig for now, fill into val beforehand
	ov := cuest.ctx.CompileBytes(origs[0].Content, cue.Filename(origs[0].Filename))
	if ov.Err() != nil {
		return nil, ov.Err()
	}
	// update val with the orig value
	val = val.FillPath(cue.ParsePath("val.#U"), ov)

	upserts := make([]UpsertResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#X"), iv)

		dv := result.LookupPath(cue.ParsePath("upsert"))

		out, err := FormatOutput(dv, flags.RootPflags.Out)
		if err != nil {
			return nil, err
		}

		upserts = append(upserts, UpsertResult{
			Filename: input.Filename,
			Content:  out,
		})

	}

	return upserts, nil
}

