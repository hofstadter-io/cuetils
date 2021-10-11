package structural

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"

	"github.com/hofstadter-io/cuetils"
)

type Cuest struct {
	op string
	ctx *cue.Context
	orig cue.Value
}

const cuemod = `
module: "github.com/hofstadter-io/cuetils"
`

func NewCuest(op string) (Cuest, error) {
	cuest := Cuest{
		op: op,
	}
	cuest.ctx = cuecontext.New()

	rd, err := cuetils.CueEmbeds.ReadFile("recurse/recurse.cue")
	if err != nil {
		return cuest, err
	}
	sf := fmt.Sprintf("structural/%s.cue", op)
	sd, err := cuetils.CueEmbeds.ReadFile(sf)
	if err != nil {
		return cuest, err
	}

	cfg := load.Config {
		ModuleRoot: "/cuetils",
		Overlay: make(map[string]load.Source),
		Dir: "/cuetils",
	}
	cfg.Overlay["/cuetils/cue.mod/module.cue"] = load.FromString(cuemod)
	cfg.Overlay["/cuetils/recurse/recuse.cue"] = load.FromBytes(rd)
	cfg.Overlay["/cuetils/" + sf] = load.FromBytes(sd)

	bis := load.Instances([]string{sf}, &cfg)

	bi := bis[0]
	// check for errors on the instance
	// these are typically parsing errors
	if bi.Err != nil {
		return cuest, bi.Err
	}

	// Use cue.Context to turn build.Instance to cue.Instance
	value := cuest.ctx.BuildInstance(bi)
	if value.Err() != nil {
		return cuest, value.Err()
	}

	// Validate the value
	err = value.Validate()
	if err != nil {
		return cuest, err
	}

	cuest.orig = value

	return cuest, nil
}
