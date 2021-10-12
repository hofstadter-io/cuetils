package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"

	"github.com/hofstadter-io/cuetils"
)

type Cuest struct {
	ctx *cue.Context
	orig cue.Value
}

const cuemod = `
module: "github.com/hofstadter-io/cuetils"
`

// Creates a new Cuest object and preloads the structural ops.
// The ops should be their lowercase name and all are loaded.
// If 'op' is nil or empty, preloading is skipped.
//
// You can provide your own cue.Context or a new one will be created
// Keep in mind CUE still requires values to come from the same context.
func NewCuest(ops []string, ctx *cue.Context) (*Cuest, error) {
	cuest := new(Cuest)
	if ctx == nil {
		cuest.ctx = cuecontext.New()
	} else {
		cuest.ctx = ctx
	}

	if ops == nil || len(ops) == 0 {
		return cuest, nil
	}

	rd, err := cuetils.CueEmbeds.ReadFile("recurse/recurse.cue")
	if err != nil {
		return nil, err
	}
	cfg := load.Config {
		ModuleRoot: "/cuetils",
		Overlay: make(map[string]load.Source),
		Dir: "/cuetils",
	}
	cfg.Overlay["/cuetils/cue.mod/module.cue"] = load.FromString(cuemod)
	cfg.Overlay["/cuetils/recurse/recurse.cue"] = load.FromBytes(rd)

	entrypoints := []string{}
	for _,op := range ops {
		sf := fmt.Sprintf("structural/%s.cue", op)
		sd, err := cuetils.CueEmbeds.ReadFile(sf)
		if err != nil {
			return nil, err
		}
		cfg.Overlay["/cuetils/" + sf] = load.FromBytes(sd)
		entrypoints = append(entrypoints, sf)
	}

	bis := load.Instances(entrypoints, &cfg)

	bi := bis[0]
	// check for errors on the instance
	// these are typically parsing errors
	if bi.Err != nil {
		return nil, bi.Err
	}

	// Use cue.Context to turn build.Instance to cue.Instance
	value := cuest.ctx.BuildInstance(bi)
	if value.Err() != nil {
		return nil, value.Err()
	}

	// Validate the value
	err = value.Validate()
	if err != nil {
		return nil, err
	}

	cuest.orig = value

	return cuest, nil
}
