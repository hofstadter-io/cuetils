package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type CountResult struct {
	Filename string
	Count int
}

func Count(globs []string, rflags flags.RootPflagpole) ([]CountResult, error) {
	// no globs, then stdin
	if len(globs) == 0 {
		globs = []string{"-"}
	}

	inputs, err := ReadGlobs(globs)
	if err != nil {
		return nil, err
	}
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no matches found")
	}

	counter := func (val cue.Value) int {
		sum := 0
		after := func (v cue.Value) {
			switch v.IncompleteKind() {
				case cue.StructKind:
					s, _ := v.Fields(defaultWalkOptions...)
					for s.Next() {
						sum += 1
					}
				case cue.ListKind:
					// nothing
				default:
					sum += 1
			}
		}

		Walk(val, nil, after)
		return sum
	}

	ctx := cuecontext.New()

	counts := make([]CountResult, 0)
	for _, input := range inputs {

		// need to handle encodings here?

		iv := ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		c := counter(iv)

		counts = append(counts, CountResult{
			Filename: input.Filename,
			Count: c,
		})

	}

	return counts, nil
}
