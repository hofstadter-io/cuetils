package structural

import (
	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

func UpsertGlobs(code string, globs []string, opts *flags.RootPflagpole) ([]GlobResult, error) {
	return BinaryOpGlobs(code, globs, opts, UpsertValue)
}

func UpsertValue(up, val cue.Value, opts *flags.RootPflagpole) (cue.Value, error) {
	r, _ := upsertValue(up, val, opts)
	return r, nil
}

func upsertValue(up, val cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	switch val.IncompleteKind() {
	case cue.StructKind:
		return upsertStruct(up, val, opts)

	case cue.ListKind:
		return upsertList(up, val, opts)

	default:
		// should already have the same label by now
		// but maybe not if target is basic and up is not
		return up, true
	}
}

func upsertStruct(up, val cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	// TODO, implement proper default helper func
	// for flags in hofmod-cli
	if opts == nil {
		opts = &flags.RootPflagpole{}
	}

	ctx := val.Context()
	result := newStruct(ctx)

	// first loop over val
	iter, _ := val.Fields(defaultWalkOptions...)
	for iter.Next() {
		s := iter.Selector()
		p := cue.MakePath(s)
		u := up.LookupPath(p)

		// check that field exists in from. Should we be checking f.Err()?
		if u.Exists() {
			r, ok := upsertValue(u, iter.Value(), opts)
			// fmt.Println("r:", r, ok, p)
			if ok {
				result = result.FillPath(p, r)
			}
		} else {
			// include if not in val
			result = result.FillPath(p, iter.Value())
		}
	}

	// add anything in ins that is not in val
	iter, _ = up.Fields(defaultWalkOptions...)
	for iter.Next() {
		s := iter.Selector()
		p := cue.MakePath(s)
		v := val.LookupPath(p)

		// check that field exists in from. Should we be checking f.Err()?
		if !v.Exists() {
			result = result.FillPath(p, iter.Value())
		}
	}

	return result, true
}

func upsertList(up, val cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	ctx := val.Context()

	ui, _ := up.List()
	vi, _ := val.List()

	result := []cue.Value{}
	for ui.Next() && vi.Next() {
		r, ok := upsertValue(ui.Value(), vi.Value(), opts)
		if ok {
			result = append(result, r)
		}
	}
	return ctx.NewList(result...), true

}
