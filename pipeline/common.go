package pipeline

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/format"
)

func formatCue(val cue.Value) (string, error) {
	syn := val.Syntax(
		cue.Final(),
		cue.ResolveReferences(true),
		cue.Concrete(true),
		cue.Definitions(true),
		cue.Hidden(true),
		cue.Optional(true),
		cue.Attributes(true),
		cue.Docs(true),
	)

	bs, err := format.Node(syn)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}
