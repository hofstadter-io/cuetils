package structural

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

const maskfmt = `
val: #Mask%s
val: #X: _
val: #M: _
mask: val.mask
`

// MaskGlobs will mask a subobject from globs on disk
func MaskGlobs(mask string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	return MaskGlobsCue(mask, globs, rflags)
}

func MaskGlobsGo(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	return BinaryOpGlobs(code, globs, rflags, MaskValue)
}

func MaskValue(mask, val cue.Value) (cue.Value, error) {
	r, _ := maskValue(mask, val)
	return r, nil
}

func maskValue(mask, val cue.Value) (cue.Value, bool) {
	return mask, false
}

func MaskGlobsCue(mask string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	cuest, err := NewCuest([]string{"mask"}, nil)
	if err != nil {
		return nil, err
	}

	operator, err := ReadArg(mask, cuest.ctx, nil)
	if err != nil {
		return nil, err
	}

	inputs, err := LoadGlobs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	// construct reusable val with function
	maxiter := ""
	if mi := rflags.Maxiter; mi > 0 {
		maxiter = fmt.Sprintf(" & { #maxiter: %d }", mi)
	}
	content := fmt.Sprintf(maskfmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	// fill val with the orig value, so we only need to once before loop
	val = val.FillPath(cue.ParsePath("val.#M"), operator.Value)

	results := make([]GlobResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#X"), iv)

		v := result.LookupPath(cue.ParsePath("mask"))

		results = append(results, GlobResult{
			Filename: input.Filename,
			Value:    v,
		})

	}

	return results, nil
}
