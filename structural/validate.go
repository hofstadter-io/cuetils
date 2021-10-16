package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

func ValidateGlobs(schema string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	cuest, err := NewCuest(nil, nil)
	if err != nil {
		return nil, err
	}

	operator, err := ParseOperator(schema)
	if err != nil {
		return nil, err
	}
	operator, err = LoadOperator(operator, rflags.Load, cuest.ctx)
	if err != nil {
		return nil, err
	}

	inputs, err := ReadGlobs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	results := make([]GlobResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
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
	e := val.Unify(schema).Err()
	return e == nil, e
}
