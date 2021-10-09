package structural

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

type Input struct {
	Filename   string
	Filetype   string  // yaml, json, cue... toml?
	Expression string  // cue expression to select within document
	Content    []byte
}

func LoadInputs(globs []string) ([]Input, error) {

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
