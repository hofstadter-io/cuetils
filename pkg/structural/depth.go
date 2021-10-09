package structural

import (
	"fmt"
	"os"
	"path/filepath"

	"cuelang.org/go/cue"
)


type FileDepth struct {
	Filename string
	Depth int
}

func Depth(globs []string) ([]FileDepth, error) {
	depths := make([]FileDepth, 0)

	matches := make([]string, 0)
	for _, g := range globs {
		ms, err := filepath.Glob(g)
		if err != nil {
			return depths, err
		}
		matches = append(matches, ms...)
	}

	if len(matches) == 0 {
		return depths, fmt.Errorf("no matches found")
	}

	cuest, err := NewCuest("depth")
	if err != nil {
		return depths, err
	}

	ctx := cuest.ctx
	val := ctx.CompileString("val: #Depth\nval: #in: _\ndepth: val.out", cue.Scope(cuest.orig))

	for _, m := range matches {
		d, err := os.ReadFile(m)
		if err != nil {
			return depths, err
		}

		input := ctx.CompileBytes(d, cue.Filename(m))
		if input.Err() != nil {
			return depths, input.Err()
		}

		result := val.FillPath(cue.ParsePath("val.#in"), input)

		dv := result.LookupPath(cue.ParsePath("depth"))
		di, err := dv.Int64()
		if err != nil {
			return depths, err
		}

		depths = append(depths, FileDepth{
			Filename: m,
			Depth: int(di),
		})

	}

	return depths, nil
}
