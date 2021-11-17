package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

const transformfmt = `
#Transformer: _
#In: _
Out: #Transformer
`

func TransformGlobs(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	cuest, err := NewCuest(nil, nil)
	if err != nil {
		return nil, err
	}

	operator, err := ReadArg(code, rflags.Load, cuest.ctx, nil)
	if err != nil {
		return nil, err
	}

	inputs, err := LoadGlobs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	val := cuest.ctx.CompileString(transformfmt)

	// fill val with the orig value, so we only need to once before loop
	val = val.FillPath(cue.ParsePath("#Transformer"), operator.Value)

	results := make([]GlobResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("#Transformer.#In"), iv)

		v := result.LookupPath(cue.ParsePath("Out"))

		results = append(results, GlobResult{
			Filename: input.Filename,
			Value:    v,
		})

	}

	return results, nil
}

func TransformValue(trans, orig cue.Value) (cue.Value, error) {

	return orig, nil
}
