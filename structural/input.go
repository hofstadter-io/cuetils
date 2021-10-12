package structural

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
)

type Input struct {
	Filename   string
	Filetype   string  // yaml, json, cue... toml?
	Expression string  // cue expression to select within document
	Content    []byte
}

// Loads the entrypoints using the context provided
// returns the value from the load after validating it
func LoadInputs(entrypoints []string, ctx *cue.Context) (cue.Value, error) {

	bis := load.Instances(entrypoints, nil)

	bi := bis[0]
	// check for errors on the instance
	// these are typically parsing errors
	if bi.Err != nil {
		return cue.Value{}, bi.Err
	}

	// Use cue.Context to turn build.Instance to cue.Instance
	value := ctx.BuildInstance(bi)
	if value.Err() != nil {
		return cue.Value{}, value.Err()
	}

	// Validate the value
	err := value.Validate(
		cue.ResolveReferences(true),
		cue.Concrete(false),
		cue.Definitions(true),
		cue.Hidden(true),
		cue.Optional(true),
		cue.Attributes(false),
		cue.Docs(false),
	)
	if err != nil {
		return cue.Value{}, err
	}

	return value, nil
}

func ReadGlobs(globs []string) ([]Input, error) {

	// handle special stdin case
	if len(globs) == 1 && globs[0] == "-" {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		i := []Input{ Input{Filename: "-",Content: b} }
		return i, nil
	}

	// handle general case
	// we will load into a map to remove duplicates
	// and then extract and sort in a slice
	inputs := make(map[string]Input)
	for _, g := range globs {
		// need to check for expression syntax here

		matches, err := filepath.Glob(g)
		if err != nil {
			return nil, err
		}

		for _, m := range matches {
			// continue on duplicate
			if _,ok := inputs[m]; ok {
				continue
			}

			d, err := os.ReadFile(m)
			if err != nil {
				return nil, err
			}

			// handle input types
			ext := filepath.Ext(m)
			switch ext {
			case ".yml", ".yaml":
				s := fmt.Sprintf(yamlMod, string(d))
				d = []byte(s)
			}

			inputs[m] = Input{ Filename: m, Content: d }
		}
	}

	ret := make([]Input, 0)
	for _, i := range inputs {
		ret = append(ret, i)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Filename < ret[j].Filename
	})

	return ret, nil
}

const yamlMod = `
import "encoding/yaml"
#content: #"""
%s
"""#
yaml.Unmarshal(#content)
`
