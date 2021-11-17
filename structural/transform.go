package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

const transformfmt = `
#Transformer: _
#In: _
Out: #Transformer
`

func TransformGlobs(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	ctx := cuecontext.New()

	ov, err := ReadArg(code, rflags.Load, ctx, nil)
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

		v, err := TransformValue(ov.Value, iv)
		if err != nil {
			return nil, err
		}

		// TODO, accumulate error in results and continue looping

		results = append(results, GlobResult{
			Filename: input.Filename,
			Value:    v,
		})

	}

	return results, nil
}

func TransformValue(trans, orig cue.Value) (cue.Value, error) {
	ctx := trans.Context()
	val := ctx.CompileString(transformfmt)
	val = val.FillPath(cue.ParsePath("#Transformer"), trans)
	val = val.FillPath(cue.ParsePath("#In"), trans)
	out := val.LookupPath(cue.ParsePath("Out"))

	return out, nil
}
