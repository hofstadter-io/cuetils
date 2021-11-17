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

func DepthGlobs(globs []string, rflags flags.RootPflagpole) ([]DepthResult, error) {
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

	depths := make([]DepthResult, 0)
	for _, input := range inputs {

		// need to handle encodings here?

		iv := ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		d := depth(iv)

		depths = append(depths, DepthResult{
			Filename: input.Filename,
			Depth:    int(d),
		})

	}

	return depths, nil
}

func DepthValue(val cue.Value) int {
	return depth(val)
}

func depth(val cue.Value) int {
	var max, curr int

	// increase curr, check against max
	before := func(v cue.Value) bool {
		switch v.IncompleteKind() {
		case cue.StructKind:
			curr += 1
		case cue.ListKind:
			// nothing
		default:
			curr += 1
		}

		if curr > max {
			max = curr
		}
		return true
	}
	// decrease curr after
	after := func(v cue.Value) {
		switch v.IncompleteKind() {
		case cue.StructKind:
			curr -= 1
		case cue.ListKind:
			// nothing
		default:
			curr -= 1
		}
	}

	Walk(val, before, after)
	return max
}
