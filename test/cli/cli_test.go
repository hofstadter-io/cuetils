package structural_test

import (
	"testing"

	"github.com/hofstadter-io/hof/lib/yagu"
	"github.com/hofstadter-io/hof/script/runtime"
)

func TestCliTests(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir: "testdata",
		Glob: "*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

//func TestCliBugs(t *testing.T) {
	//yagu.Mkdir(".workdir/bugs")
	//runtime.Run(t, runtime.Params{
		//Dir: "testdata/bugs",
		//Glob: "*.txt",
		//WorkdirRoot: ".workdir/bugs",
	//})
//}
