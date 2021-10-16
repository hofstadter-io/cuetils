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

func TransformGlobs(transformer string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	cuest, err := NewCuest(nil, nil)
	if err != nil {
		return nil, err
	}

	ov, err := LoadInputs([]string{transformer}, cuest.ctx)
	if err != nil {
		return nil, err
	}
	if ov.Err() != nil {
		return nil, ov.Err()
	}

	inputs, err := ReadGlobs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	val := cuest.ctx.CompileString(transformfmt)

	// fill val with the orig value, so we only need to once before loop
	val = val.FillPath(cue.ParsePath("#Transformer"), ov)

	results := make([]GlobResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("#Transformer.#In"), iv)

		dv := result.LookupPath(cue.ParsePath("Out"))

		out, err := FormatOutput(dv, rflags.Out)
		if err != nil {
			return nil, err
		}

		results = append(results, GlobResult{
			Filename: input.Filename,
			Content:  out,
		})

	}

	return results, nil
}


func TransformValue(trans, orig cue.Value) (cue.Value, error) {

	return orig, nil
}
