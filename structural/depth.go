package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type DepthResult struct {
	Filename string
	Depth int
}

const depthfmt = `
val: #Depth%s
val: #in: _
depth: val.depth
`

func Depth(globs []string) ([]DepthResult, error) {
	// no globs, then stdin
	if len(globs) == 0 {
		globs = []string{"-"}
	}

	inputs, err := LoadInputs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no matches found")
	}

	cuest, err := NewCuest("depth")
	if err != nil {
		return nil, err
	}

	// construct reusable val with function
	maxiter := ""
	if mi := flags.RootPflags.Maxiter; mi > 0 {
		maxiter = fmt.Sprintf(" & { #maxiter: %d }", mi)
	}
	content := fmt.Sprintf(depthfmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	depths := make([]DepthResult, 0)
	for _, input := range inputs {

		// need to handle encodings here

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#in"), iv)

		dv := result.LookupPath(cue.ParsePath("depth"))
		di, err := dv.Int64()
		if err != nil {
			return nil, err
		}

		depths = append(depths, DepthResult{
			Filename: input.Filename,
			Depth: int(di),
		})

	}

	return depths, nil
}
