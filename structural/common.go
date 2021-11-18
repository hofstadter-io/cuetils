package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

func newAny(ctx *cue.Context) cue.Value {
	return ctx.CompileString("_")
}

func newStruct(ctx *cue.Context) cue.Value {
	return ctx.CompileString("{...}")
}

func GetLabel(val cue.Value) cue.Selector {
	ss := val.Path().Selectors()
	s := ss[len(ss)-1]
	return s
}

type BinaryOpValueFunc func(lhs, rhs cue.Value) (cue.Value, error)

func BinaryOpGlobs(lhs string, rhs []string, rflags flags.RootPflagpole, fn BinaryOpValueFunc) ([]GlobResult, error) {
	ctx := cuecontext.New()

	lv, err := ReadArg(lhs, rflags.Load, ctx, nil)
	if err != nil {
		return nil, err
	}

	vals, err := LoadGlobs(rhs)
	if len(vals) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	results := make([]GlobResult, 0)
	for _, val := range vals {

		iv := ctx.CompileBytes(val.Content, cue.Filename(val.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		// Call our OpValueFunc
		v, err := fn(lv.Value, iv)
		if err != nil {
			return nil, err
		}

		results = append(results, GlobResult{
			Filename: val.Filename,
			Value:    v,
		})
	}

	return results, nil
}

type UnaryOpValueFunc func(val cue.Value) (cue.Value, error)

func UnaryOpGlobs(globs []string, rflags flags.RootPflagpole, fn UnaryOpValueFunc) ([]GlobResult, error) {
	ctx := cuecontext.New()

	vals, err := LoadGlobs(globs)
	if err != nil {
		return nil, err
	}
	if len(vals) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	results := make([]GlobResult, 0)
	for _, val := range vals {

		iv := ctx.CompileBytes(val.Content, cue.Filename(val.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		// Call our OpValueFunc
		v, err := fn(iv)
		if err != nil {
			return nil, err
		}

		results = append(results, GlobResult{
			Filename: val.Filename,
			Value:    v,
		})
	}

	return results, nil
}
