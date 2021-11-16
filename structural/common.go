package structural

import (
	"fmt"

	"cuelang.org/go/cue"
)

func newStruct(ctx *cue.Context) cue.Value {
	return ctx.CompileString("_")
	// return ctx.BuildExpr(ast.NewIdent("_"))
}

func getLabel(val cue.Value) {
	ss := val.Path().Selectors()
	s := ss[len(ss)-1]
	fmt.Println(ss, s)
}
