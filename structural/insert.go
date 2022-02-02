package structural

import (
	"cuelang.org/go/cue"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

func InsertGlobs(code string, globs []string, opts *flags.RootPflagpole) ([]GlobResult, error) {
	return BinaryOpGlobs(code, globs, opts, InsertValue)
}

func InsertValue(ins, val cue.Value, opts *flags.RootPflagpole) (cue.Value, error) {
	r, _ := insertValue(ins, val, opts)
	return r, nil
}

func insertValue(ins, val cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	// TODO, implement proper default helper func
	// for flags in hofmod-cli
	if opts == nil {
		opts = &flags.RootPflagpole{}
	}

	switch val.IncompleteKind() {
	case cue.StructKind:
		return insertStruct(ins, val, opts)

	case cue.ListKind:
		return insertList(ins, val, opts)

	default:
		// should already have the same label by now
		// but maybe not if target is basic and repl is not
		return val, true
	}
}

func insertStruct(ins, val cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {

	result := val
	iter, _ := ins.Fields(defaultWalkOptions...)

	for iter.Next() {
		s := iter.Selector()
		p := cue.MakePath(s)
		v := val.LookupPath(p)

		// check that field exists in from. Should we be checking f.Err()?
		if v.Exists() {
			r, ok := insertValue(iter.Value(), v, opts)
			// fmt.Println("r:", r, ok, p)
			if ok {
				result = result.FillPath(p, r)
			}
		} else {
			// include if not in val
			result = result.FillPath(p, iter.Value())
		}
	}

	return result, true
}

func insertList(ins, val cue.Value, opts *flags.RootPflagpole) (cue.Value, bool) {
	ctx := val.Context()

	ii, _ := ins.List()
	vi, _ := val.List()

	result := []cue.Value{}
	for ii.Next() && vi.Next() {
		r, ok := insertValue(ii.Value(), vi.Value(), opts)
		if ok {
			result = append(result, r)
		}
	}
	return ctx.NewList(result...), true
}
