package structural

import (
	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

const transformfmt = `
#Transformer: _
Out: #Transformer
`

func TransformGlobs(code string, globs []string, opts *flags.RootPflagpole) ([]GlobResult, error) {
	return BinaryOpGlobs(code, globs, opts, TransformValue)
}

func TransformValue(trans, orig cue.Value, opts *flags.RootPflagpole) (cue.Value, error) {
	ctx := trans.Context()
	val := ctx.CompileString(transformfmt)
	val = val.FillPath(cue.ParsePath("#Transformer"), trans)
	val = val.FillPath(cue.ParsePath("#Transformer.#In"), orig)
	out := val.LookupPath(cue.ParsePath("Out"))

	return out, nil
}
