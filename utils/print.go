package utils

import (
	"fmt"

	"cuelang.org/go/cue"
)

func PrintAttr(attr cue.Attribute, val cue.Value) error {
	// maybe print
	if attr.Err() == nil {
		fmt.Println("PrintAttr", attr)
		for i := 0; i < attr.NumArgs(); i++ {
			a, _ := attr.String(i)
			v := val.LookupPath(cue.ParsePath(a))
			s, err := FormatCue(v)
			if err != nil {
				fmt.Println("Fmt error", err)
			}
			fmt.Printf("%s: %v\n", a, s)
		}
	}

	return nil
}
