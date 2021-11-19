package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

func ValidateGlobs(schema string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	ctx := cuecontext.New()

	operator, err := ReadArg(schema, ctx, nil)
	if err != nil {
		return nil, err
	}

	inputs, err := LoadGlobs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	results := make([]GlobResult, 0)
	for _, input := range inputs {

		iv := ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := iv.Unify(operator.Value)
		err := result.Validate()
		if err != nil {
			out := FormatCueError(err)
			results = append(results, GlobResult{
				Filename: input.Filename,
				Content:  out,
			})

			continue
		}
	}

	return results, nil

}

func ValidateValue(schema, val cue.Value) (bool, error) {
	// probably need to deal with some flags here...
	e := val.Unify(schema).Err()
	return e == nil, e
}
