package structural_test

import (
	"testing"

	"github.com/hofstadter-io/hof/lib/yagu"
	"github.com/hofstadter-io/hof/script/runtime"
)

func TestCliCount(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "count_*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

func TestCliDepth(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "depth_*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

func TestCliDiff(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "diff_*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

func TestCliMask(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "mask_*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

func TestCliPatch(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "patch_*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

func TestCliPick(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "pick_*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

func TestCliReplace(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "replace_*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

func TestCliTransform(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "transform_*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

func TestCliInsert(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "insert_*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

func TestCliUpsert(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "upsert_*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

func TestCliValidate(t *testing.T) {
	yagu.Mkdir(".workdir/tests")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata",
		Glob:        "validate_*.txt",
		WorkdirRoot: ".workdir/tests",
	})
}

func TestCliBugs(t *testing.T) {
	yagu.Mkdir(".workdir/bugs")
	runtime.Run(t, runtime.Params{
		Dir:         "testdata/bugs",
		Glob:        "*.txt",
		WorkdirRoot: ".workdir/bugs",
	})
}
