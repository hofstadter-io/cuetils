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

type listProc int

const (
	LIST_ERR listProc = iota // unknown arg
	LIST_AND                 // all apply
	LIST_OR                  // any apply
	LIST_PER                 // pairwise
)

func getListProcType(val cue.Value) (listProc, error) {
	attr := val.Attribute("list")
	// no attribute or no arg, so default is LIST_AND
	if attr.Err() != nil || attr.NumArgs() == 0 {
		return LIST_AND, nil
	}
	a, err := attr.String(0)
	if err != nil {
		return LIST_ERR, err
	}
	// otherwise, check what arg we have
	switch a {
	case "and", "":
		return LIST_AND, nil
	case "or":
		return LIST_OR, nil
	case "pairwise", "per":
		return LIST_PER, nil
	default:
		return LIST_ERR, fmt.Errorf("Unknown list processing type %q at %v", a, val.Pos())
	}
}

func GetLabel(val cue.Value) cue.Selector {
	ss := val.Path().Selectors()
	s := ss[len(ss)-1]
	return s
}

type BinaryOpValueFunc func(lhs, rhs cue.Value, opts *flags.RootPflagpole) (cue.Value, error)

func BinaryOpGlobs(lhs string, rhs []string, opts *flags.RootPflagpole, fn BinaryOpValueFunc) ([]GlobResult, error) {
	ctx := cuecontext.New()

	lv, err := ReadArg(lhs, ctx, nil)
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
		v, err := fn(lv.Value, iv, opts)
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

type UnaryOpValueFunc func(val cue.Value, opts *flags.RootPflagpole) (cue.Value, error)

func UnaryOpGlobs(globs []string, opts *flags.RootPflagpole, fn UnaryOpValueFunc) ([]GlobResult, error) {
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
		v, err := fn(iv, opts)
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
