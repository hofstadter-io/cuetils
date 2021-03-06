package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type StatsResult struct {
	Filename string
	Count    int
	Depth    int
}

type CountResult struct {
	Filename string
	Count    int
}

func CountGlobs(globs []string, opts *flags.RootPflagpole) ([]CountResult, error) {
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

	ctx := cuecontext.New()

	counts := make([]CountResult, 0)
	for _, input := range inputs {

		iv := ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		c := counter(iv)

		counts = append(counts, CountResult{
			Filename: input.Filename,
			Count:    c,
		})

	}

	return counts, nil
}

func CountValue(val cue.Value) int {
	return counter(val)
}

func counter(val cue.Value) int {
	sum := 0
	after := func(v cue.Value) {
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
