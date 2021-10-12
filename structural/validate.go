package structural

import (
	"cuelang.org/go/cue"
)


func ValidateFiles(schema string, globs []string) {

}

func ValidateValues(schema, val cue.Value) (error) {
	return val.Unify(schema).Err()
}
