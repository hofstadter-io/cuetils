package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
	"github.com/hofstadter-io/cuetils/structural"
)

var patchLong = `apply the pacth to the glob file(s)`

func PatchRun(patch string, globs []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	results, err := structural.Patch(patch, globs, flags.RootPflags)
	if err != nil {
		return err
	}

	err = structural.ProcessOutputs(results, flags.RootPflags)

	return err
}

var PatchCmd = &cobra.Command{

	Use: "patch <patch> <glob>",

	Short: "apply the pacth to the glob file(s)",

	Long: patchLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'patch'")
			cmd.Usage()
			os.Exit(1)
		}

		var patch string

		if 0 < len(args) {

			patch = args[0]

		}

		var globs []string

		if 1 < len(args) {

			globs = args[1:]

		}

		err = PatchRun(patch, globs)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	extra := func(cmd *cobra.Command) bool {

		return false
	}

	ohelp := PatchCmd.HelpFunc()
	ousage := PatchCmd.UsageFunc()
	help := func(cmd *cobra.Command, args []string) {
		if extra(cmd) {
			return
		}
		ohelp(cmd, args)
	}
	usage := func(cmd *cobra.Command) error {
		if extra(cmd) {
			return nil
		}
		return ousage(cmd)
	}

	PatchCmd.SetHelpFunc(help)
	PatchCmd.SetUsageFunc(usage)

}
