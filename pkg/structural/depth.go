package structural

import (
	"fmt"

	"cuelang.org/go/cue"
)


type FileDepth struct {
	Filename string
	Depth int
}

func Depth(globs []string) ([]FileDepth, error) {
	// no globs, then stdin
	if len(globs) == 0 {
		globs = []string{"-"}
	}

	inputs, err := LoadInputs(globs)
	if len(inputs) == 0 {
		return nil, fmt.Errorf("no matches found")
	}

	cuest, err := NewCuest("depth")
	if err != nil {
		return nil, err
	}

	val := cuest.ctx.CompileString("val: #Depth\nval: #in: _\ndepth: val.out", cue.Scope(cuest.orig))

	depths := make([]FileDepth, 0)

	for _, input := range inputs {

		// need to handle encodings here

		iv := cuest.ctx.CompileBytes(input.Content, cue.Filename(input.Filename))
		if iv.Err() != nil {
			return nil, iv.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#in"), iv)

		dv := result.LookupPath(cue.ParsePath("depth"))
		di, err := dv.Int64()
		if err != nil {
			return nil, err
		}

		depths = append(depths, FileDepth{
			Filename: input.Filename,
			Depth: int(di),
		})

	}

	return depths, nil
}
