package structural_test

import (
	"fmt"
	"strings"
	"testing"

	"cuelang.org/go/cue/cuecontext"
	"github.com/hofstadter-io/cuetils/structural"
)

func TestPick(t *testing.T) {
	type test struct {
		pick string
		from string
		want string
	}

	tests := []test{
		{pick: "_", from: "a: b: int", want: `
{
	a: {
		b: int
	}
}
		`},
	}

	work := []test{
		{pick: "a: _", from: "a: b: int", want: `
{
	a: {
		b: int
	}
}
		`},
	}
	fmt.Println(len(tests))

	for i, tc := range work {
		ctx := cuecontext.New()
		p := ctx.CompileString(tc.pick)
		if p.Err() != nil {
			t.Fatalf("failed to compile pick in TestPick %d", i)
		}
		f := ctx.CompileString(tc.from)
		if f.Err() != nil {
			t.Fatalf("failed to compile from in TestPick %d", i)
		}
		r, _ := structural.PickValue(p, f)

		o, err := structural.FormatOutput(r, "cue")
		if err != nil {
			t.Fatalf("failed to format output in TestPick %d", i)
		}

		o = strings.TrimSpace(o)
		tc.want = strings.TrimSpace(tc.want)

		if o != tc.want {
			t.Fatalf("expected:\n%s\n\ngot:\n%s\n", tc.want, o)
		}
	}

}
