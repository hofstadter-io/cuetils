package structural

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/errors"

	// "cuelang.org/go/cue/errors"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

// PickGlobs will pick a subobject from globs on disk
func PickGlobs(code string, globs []string, rflags flags.RootPflagpole) ([]GlobResult, error) {
	return BinaryOpGlobs(code, globs, rflags, PickValue)
}

// PickValue uses 'pick' to pick a subvalue from 'from'
// by checking if values unify
func PickValue(pick, from cue.Value) (cue.Value, error) {
	p, _ := pickValue(pick, from)
	return p, nil
}

// this is the recursive version that also returns
// whether the value was picked
func pickValue(pick, from cue.Value) (cue.Value, bool) {
	ctx := pick.Context()
	//fmt.Println(pick)
	//fmt.Println(from)

	switch pick.IncompleteKind() {
	// just return
	case cue.TopKind:
		return from, true

	// recurse on matching labels
	case cue.StructKind:
		if k := from.IncompleteKind(); k != cue.StructKind {
			// should this return or just continue? do we need some way to specify?
			// probably prefer to be more strict, so that you know your schemas
			// return errors.Newf(from.Pos(), "expected struct, but got %v", k), true
			return newStruct(ctx), false
		}
		result := newStruct(ctx)
		iter, _ := pick.Fields(defaultWalkOptions...)

		cnt := 0
		for iter.Next() {
			cnt++
			s := iter.Selector()
			p := cue.MakePath(s)
			f := from.LookupPath(p)
			// fmt.Println(cnt, iter.Value(), f, f.Exists())
			// check that field exists in from. Should we be checking f.Err()?
			if f.Exists() {
				r, ok := pickValue(iter.Value(), f)
				// fmt.Println("r:", r, ok, p)
				if ok {
					result = result.FillPath(p, r)
				}
			}
		}

		// need to check for {...}
		// no fields and open
		if cnt == 0 && pick.Allows(cue.AnyString) {
			return from, true
		}

		// fmt.Println("result:", result)

		return result, true

	case cue.ListKind:
		if k := from.IncompleteKind(); k != cue.ListKind {
			// should this return or just continue? do we need some way to specify?
			// probably prefer to be more strict, so that you know your schemas
			// return errors.Newf(from.Pos(), "expected list, but got %v", k), true
			return newStruct(ctx), false
		}

		lpt, err := getListProcType(pick)
		if err != nil {
			ce := errors.Newf(pick.Pos(), "%v", err)
			ev := ctx.MakeError(ce)
			return ev, true
		}

		_ = lpt

		// how to consider different list sizes
		// if len(pick) == 1, apply to all elements
		// if len(pick) > 1
		//   attributes? @pick(and,or,pos)
		// maybe we don't care about length if attribute is used?

		pi, _ := pick.List()
		fi, _ := from.List()

		result := []cue.Value{}
		for pi.Next() && fi.Next() {
			p, ok := pickValue(pi.Value(), fi.Value())
			if ok {
				result = append(result, p)
			}
		}

		return ctx.NewList(result...), true

	// (basic lit types)
	default:
		// if pick is concrete, so must from
		// make sure 1 does not pick int
		// but we do want int to pick any num
		if pick.IsConcrete() {
			if from.IsConcrete() {
				r := pick.Unify(from)
				return r, r.Exists()
			} else {
				return cue.Value{}, false
			}
		} else {
			r := pick.Unify(from)
			return r, r.Exists()
		}

	}
}
