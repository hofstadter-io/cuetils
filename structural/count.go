package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"

	// "github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type CountResult struct {
	Filename string
	Count int
}

func Count(globs []string) ([]CountResult, error) {
	// no globs, then stdin
	if len(globs) == 0 {
		globs = []string{"-"}
	}

	//cuest, err := NewCuest([]string{"count"}, nil)
	//if err != nil {
		//return nil, err
	//}

	inputs, err := ReadGlobs(globs)
	if err != nil {
		return nil, err
	}
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no matches found")
	}

	counter := func (val cue.Value) int {
		sum := 0
		after := func (v cue.Value) {
			switch v.IncompleteKind() {
				case cue.StructKind:
					sum += 1
				case cue.ListKind:
					// nothing
				default:
					sum += 1
			}
		}

		Walk(val, nil, after)
		return sum
	}

	// construct reusable val with function
	//maxiter := ""
	//if mi := flags.RootPflags.Maxiter; mi > 0 {
		//maxiter = fmt.Sprintf(" & { #maxiter: %d }", mi)
	//}
	//content := fmt.Sprintf(countfmt, maxiter)
	//val := cuest.ctx.CompileString(content, cue.Scope(cuest.orig))

	ctx := cuecontext.New()

	counts := make([]CountResult, 0)
	for _, input := range inputs {

		// need to handle encodings here

		iv := ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		c := counter(iv)

		//result := val.FillPath(cue.ParsePath("val.#in"), iv)

		//dv := result.LookupPath(cue.ParsePath("count"))
		//di, err := dv.Int64()
		//if err != nil {
			//return nil, err
		//}

		counts = append(counts, CountResult{
			Filename: input.Filename,
			Count: c,
		})

	}

	return counts, nil
}
