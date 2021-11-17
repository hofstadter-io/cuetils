package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type DepthResult struct {
	Filename string
	Depth    int
}

func Depth(globs []string, rflags flags.RootPflagpole) ([]DepthResult, error) {
	// no globs, then stdin
	if len(globs) == 0 {
		globs = []string{"-"}
	}

	inputs, err := LoadGlobs(globs)
	if err != nil {
		return nil, err
	}
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no matches found")
	}

	depther := func(val cue.Value) int {
		var max, depth int

		// increase depth, check against max
		before := func(v cue.Value) bool {
			switch v.IncompleteKind() {
			case cue.StructKind:
				depth += 1
			case cue.ListKind:
				// nothing
			default:
				depth += 1
			}

			if depth > max {
				max = depth
			}
			return true
		}
		// decrease depth after
		after := func(v cue.Value) {
			switch v.IncompleteKind() {
			case cue.StructKind:
				depth -= 1
			case cue.ListKind:
				// nothing
			default:
				depth -= 1
			}
		}

		Walk(val, before, after)
		return max
	}

	ctx := cuecontext.New()

	depths := make([]DepthResult, 0)
	for _, input := range inputs {

		// need to handle encodings here?

		iv := ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		d := depther(iv)

		depths = append(depths, DepthResult{
			Filename: input.Filename,
			Depth:    int(d),
		})

	}

	return depths, nil
}
