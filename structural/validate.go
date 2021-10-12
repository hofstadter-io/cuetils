package structural

import (
	"cuelang.org/go/cue"
)


func ValidateFiles(schema string, globs []string) {

}

func ValidateValue(schema, val cue.Value) (bool, error) {
	e := val.Unify(schema).Err()
	return e == nil, e
}
