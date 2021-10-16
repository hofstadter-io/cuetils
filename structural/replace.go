package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type ReplaceResult struct {
	Filename string
	Content  string
}

const replacefmt = `
val: #Replace%s
val: #X: _
val: #R: _
replace: val.replace
`

func Replace(orig string, globs []string, rflags flags.RootPflagpole) ([]ReplaceResult, error) {
	return ReplaceGlobsCUE(orig, globs, rflags)
}

func ReplaceGlobsCUE(orig string, globs []string, rflags flags.RootPflagpole) ([]ReplaceResult, error) {
	cuest, err := NewCuest([]string{"replace"}, nil)
	if err != nil {
		return nil, err
	}

	ov, err := LoadInputs([]string{orig}, cuest.ctx)
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

	// construct reusable val with function
	maxiter := ""
	if mi := rflags.Maxiter; mi > 0 {
		maxiter = fmt.Sprintf(" & { #maxiter: %d }", mi)
	}
	content := fmt.Sprintf(replacefmt, maxiter)
	val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	// fill val with the orig value, so we only need to once before loop
	val = val.FillPath(cue.ParsePath("val.#R"), ov)

	replaces := make([]ReplaceResult, 0)
	for _, input := range inputs {

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#X"), iv)

		dv := result.LookupPath(cue.ParsePath("replace"))

		out, err := FormatOutput(dv, rflags.Out)
		if err != nil {
			return nil, err
		}

		replaces = append(replaces, ReplaceResult{
			Filename: input.Filename,
			Content:  out,
		})

	}

	return replaces, nil
}

func ReplaceGlobsGo(orig string, globs []string, rflags flags.RootPflagpole) ([]ReplaceResult, error) {
	ctx := cuecontext.New()

	//ov, err := LoadInputs([]string{orig}, ctx)
	//if err != nil {
		//return nil, err
	//}
	//if ov.Err() != nil {
		//return nil, ov.Err()
	//}

	origs, err := ReadGlobs([]string{orig})
	if err != nil {
		return nil, err
	}
	if len(origs) == 0 {
		return nil, fmt.Errorf("no orig found")
	}
	ov := ctx.CompileBytes(origs[0].Content)

	inputs, err := ReadGlobs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no inputs found")
	}

	replaces := make([]ReplaceResult, 0)
	for _, input := range inputs {

		iv := ctx.CompileBytes(input.Content, cue.Filename(input.Filename), cue.Scope(ov))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		dv := ReplaceValue(ov, iv, ctx)

		out, err := FormatOutput(dv, flags.RootPflags.Out)
		if err != nil {
			return nil, err
		}

		replaces = append(replaces, ReplaceResult{
			Filename: input.Filename,
			Content:  out,
		})

	}

	return replaces, nil
}

func ReplaceValue(replace, orig cue.Value, ctx *cue.Context) cue.Value {
	return replacer(replace, orig, ctx)
}

func replacer(replace, orig cue.Value, ctx *cue.Context) cue.Value {
	var result cue.Value

	switch orig.IncompleteKind() {
	case cue.StructKind:
		result = ctx.CompileString("")
		p := orig.Path()
		fmt.Println("p", p)
		_, rp := orig.ReferencePath()
		fmt.Println("rp", rp)
		s, _ := orig.Fields(defaultWalkOptions...)
		for s.Next() {
			sel := s.Selector()
			fmt.Println(sel)
			path := cue.MakePath(sel)
			r := replace.LookupPath(path)
			if r.Exists() {
				tmp := replacer(s.Value(), r, ctx)
				result = result.FillPath(path, tmp)
			}
		}

	case cue.ListKind:
		vals := []cue.Value{}
		l, _ := orig.List()
		for l.Next() {
		}
		result = ctx.NewList(vals...)

	default:
		p := orig.Path()
		r := replace.LookupPath(p)
		if r.Exists() {
			result = r
		}
	}



	return result
}
