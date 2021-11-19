package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
	"github.com/hofstadter-io/cuetils/structural"
)

var diffLong = `calculate the diff from orig to next file(s)`

func DiffRun(orig string, next string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	results, err := structural.DiffGlobs(orig, next, &flags.RootPflags)
	if err != nil {
		return err
	}

	err = structural.ProcessOutputs(results, &flags.RootPflags)

	return err
}

var DiffCmd = &cobra.Command{

	Use: "diff <orig> <next>",

	Aliases: []string{
		"D",
	},

	Short: "calculate the diff from orig to next file(s)",

	Long: diffLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'orig'")
			cmd.Usage()
			os.Exit(1)
		}

		var orig string

		if 0 < len(args) {

			orig = args[0]

		}

		if 1 >= len(args) {
			fmt.Println("missing required argument: 'next'")
			cmd.Usage()
			os.Exit(1)
		}

		var next string

		if 1 < len(args) {

			next = args[1]

		}

		err = DiffRun(orig, next)
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

	ohelp := DiffCmd.HelpFunc()
	ousage := DiffCmd.UsageFunc()
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

	DiffCmd.SetHelpFunc(help)
	DiffCmd.SetUsageFunc(usage)

}
