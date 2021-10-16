package structural

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/format"
	"cuelang.org/go/encoding/yaml"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

type GlobResult struct {
	Filename string
	Content  string
	Value    cue.Value
}

func ProcessOutputs(results []GlobResult, rflags flags.RootPflagpole) (err error) {
	//if rflags.Accum != "" {
		//results, err = AccumOutputs(results, rflags.Accum)
		//if err != nil {
			//return err
		//}
	//}
	w := os.Stdout
	for _, r := range results {
		// Format
		r.Content, err = FormatOutput(r.Value, rflags.Out)
		if err != nil {
			return err
		}

		// make outname
		outname := ""
		if rflags.Outname != "" {
			outname = rflags.Outname
			// look for interpolation syntax
			if strings.Contains(outname, "<") {
				dir, file := filepath.Split(r.Filename)
				ext := filepath.Ext(file)
				name := strings.TrimSuffix(file, ext)

				outname = strings.Replace(outname, "<dir>", dir, -1)
				outname = strings.Replace(outname, "<name>", name, -1)
				outname = strings.Replace(outname, "<ext>", ext, -1)
				outname = strings.Replace(outname, "<filename>", file, -1)
				outname = strings.Replace(outname, "<filepath>", r.Filename, -1)
			}
			if strings.Contains(outname, "\\(") {
				o := r.Value.Context().CompileString(outname, cue.Scope(r.Value))
				outname, err = o.String()
				if err != nil {
					return err
				}
			}
		}

		// are we writing a file?
		writeFile := false
		if rflags.Overwrite || outname != "" {
			writeFile = true
		}
		// now possibly fill filename
		if outname == "" {
			outname = r.Filename
		}

		// if yes, we need to override w
		if writeFile {
			_, err = os.Stat(outname)
			// if no overwrite and exists, return err
			if !rflags.Overwrite && err == nil {
				return fmt.Errorf("output file %q exists, use --overwrite to replace", outname)
			}
			w, err = os.Create(outname)
			if err != nil {
				return err
			}
		}

		// now do actual writing
		if rflags.Headers {
			fmt.Fprintf(w, "%s\n----------------------\n%s\n\n", outname, r.Content)
		} else {
			fmt.Fprintf(w, "%s\n", r.Content)
		}
	}

	return nil
}

func AccumOutputs(results []GlobResult, accum string) ([]GlobResult, error) {
	return results, nil
}

func FormatOutput(val cue.Value, format string) (string, error) {
	switch format {
	case "cue", "CUE":
		return formatCue(val)

	case "json":
		return formatJson(val)

	case "yml", "yaml":
		return formatYaml(val)

	default:
		return "", fmt.Errorf("unknown output encoding %q", format)
	}

}

func formatCue(val cue.Value) (string, error) {
	syn := val.Syntax(
		cue.Final(),
		cue.ResolveReferences(true),
		cue.Concrete(true),
		cue.Definitions(false),
		cue.Hidden(false),
		cue.Optional(false),
		cue.Attributes(false),
		cue.Docs(false),
	)

	bs, err := format.Node(syn)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func formatJson(val cue.Value) (string, error) {
	var w bytes.Buffer
	d := json.NewEncoder(&w)
	d.SetIndent("", "  ")

	err := d.Encode(val)
	if _, ok := err.(*json.MarshalerError); ok {
		return "", err
	}

	return w.String(), nil
}

func formatYaml(val cue.Value) (string, error) {
	bs, err := yaml.Encode(val)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}
