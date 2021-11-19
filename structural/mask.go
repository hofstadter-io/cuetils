package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/errors"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

const maskfmt = `
val: #Mask%s
val: #X: _
val: #M: _
mask: val.mask
`

// MaskGlobs will mask a subobject from globs on disk
func MaskGlobs(mask string, globs []string, opts *flags.RootPflagpole) ([]GlobResult, error) {
	return MaskGlobsCue(mask, globs, opts)
}

func MaskGlobsGo(code string, globs []string, opts *flags.RootPflagpole) ([]GlobResult, error) {
	return BinaryOpGlobs(code, globs, opts, MaskValue)
}

func MaskValue(mask, val cue.Value, opts *flags.RootPflagpole) (cue.Value, error) {
	r, _ := maskValue(mask, val, opts)
	return r, nil
}

// returns a value and if it should be kept / part of the return
func maskValue(mask, from cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {

	switch mask.IncompleteKind() {
	// mask everything
	case cue.TopKind:
		return newStruct(from.Context()), false

	// recurse on matching labels
	case cue.StructKind:
		return maskStruct(mask, from, opts)

	case cue.ListKind:
		return maskList(mask, from, opts)
	// (basic lit types)
	default:
		return maskLeaf(mask, from, opts)
	}

}

func maskStruct(mask, from cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	ctx := mask.Context()
	eflags := opts.AllTypeErrors || opts.NodeTypeErrors

	if eflags && mask.IncompleteKind() != from.IncompleteKind() {
		// emsg, hasErrMsg := getErrorAttrMsg(mask)
		e := errors.Newf(mask.Pos(), "mask type '%v' does not match target value type '%v'", mask.IncompleteKind(), from.IncompleteKind())
		ev := ctx.MakeError(e)
		return ev, true
	}

	result := newStruct(ctx)
	iter, _ := from.Fields(defaultWalkOptions...)

	cnt := 0
	for iter.Next() {
		cnt++
		s := iter.Selector()
		p := cue.MakePath(s)
		f := from.LookupPath(p)
		// fmt.Println(cnt, iter.Value(), f, f.Exists())
		// check that field exists in from. Should we be checking f.Err()?
		if f.Exists() {
			r, keep := maskValue(iter.Value(), f, opts)
			// fmt.Println("r:", r, ok, p)
			if keep {
				result = result.FillPath(p, r)
			}
		}
	}

	// need to check for {...}
	// no fields and open
	if cnt == 0 && mask.Allows(cue.AnyString) {
		return from, true
	}

	// fmt.Println("result:", result)

	return result, true

}

func maskList(mask, from cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	ctx := mask.Context()

	if k := from.IncompleteKind(); k != cue.ListKind {
		// should this return or just continue? do we need some way to specify?
		// probably prefer to be more strict, so that you know your schemas
		// return errors.Newf(from.Pos(), "expected list, but got %v", k), true
		return newStruct(ctx), false
	}

	lpt, err := getListProcType(mask)
	if err != nil {
		ce := errors.Newf(mask.Pos(), "%v", err)
		ev := ctx.MakeError(ce)
		return ev, true
	}

	_ = lpt

	mi, _ := mask.List()
	fi, _ := from.List()

	result := []cue.Value{}
	for mi.Next() && fi.Next() {
		p, ok := maskValue(mi.Value(), fi.Value(), opts)
		if ok {
			result = append(result, p)
		}
	}

	return ctx.NewList(result...), true
}

func maskLeaf(mask, from cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	// if mask is concrete, so must from
	// make sure 1 does not mask int
	// but we do want int to mask any num

	//mc := mask.IsConcrete()
	//mt := mask.IncompleteType()
	//fc := from.IsConcrete()
	//ft := from.IncompleteType()

	//ate := opts.AllTypeErrors
	//lte := opts.LeafTypeErrors

	// both should be concrete
	//if mc && !fc {

	//}

	// ...
	//if wantsError {
	//if opts.NodeTypeErrors {
	//// emsg, hasErrMsg := getErrorAttrMsg(mask)
	//e := errors.Newf(mask.Pos(), "mask type '%v' does not match target value type '%v'", mask.IncompeteKind(), from.IncompleteKind())
	//ev := ctx.MakeError(e)
	//return ev, true
	//}
	//}

	if mask.IsConcrete() {
		if from.IsConcrete() {
			r := mask.Unify(from)
			return r, r.Exists()
		} else {
			return cue.Value{}, false
		}
	} else {
		r := mask.Unify(from)
		return r, r.Exists()
	}

}

func MaskGlobsCue(mask string, globs []string, opts *flags.RootPflagpole) ([]GlobResult, error) {
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
	if mi := opts.Maxiter; mi > 0 {
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
