package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type CountResult struct {
	Filename string
	Count int
}

const countfmt = `
val: #Count%s
val: #in: _
count: val.count
`

func Count(globs []string) ([]CountResult, error) {
	// no globs, then stdin
	if len(globs) == 0 {
		globs = []string{"-"}
	}

	cuest, err := NewCuest([]string{"count"}, nil)
	if err != nil {
		return nil, err
	}

	inputs, err := ReadGlobs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no matches found")
	}

	// construct reusable val with function
	maxiter := ""
	if mi := flags.RootPflags.Maxiter; mi > 0 {
		maxiter = fmt.Sprintf(" & { #maxiter: %d }", mi)
	}
	content := fmt.Sprintf(countfmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	counts := make([]CountResult, 0)
	for _, input := range inputs {

		// need to handle encodings here

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#in"), iv)

		dv := result.LookupPath(cue.ParsePath("count"))
		di, err := dv.Int64()
		if err != nil {
			return nil, err
		}

		counts = append(counts, CountResult{
			Filename: input.Filename,
			Count: int(di),
		})

	}

	return counts, nil
}
