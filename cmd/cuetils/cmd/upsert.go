package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
	"github.com/hofstadter-io/cuetils/structural"
)

var upsertLong = `upsert file(s) with code (extend and replace)`

func UpsertRun(code string, globs []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	results, err := structural.UpsertGlobs(code, globs, flags.RootPflags)
	if err != nil {
		return err
	}

	err = structural.ProcessOutputs(results, flags.RootPflags)

	return err
}

var UpsertCmd = &cobra.Command{

	Use: "upsert <code> [files...]",

	Aliases: []string{
		"u",
	},

	Short: "upsert file(s) with code (extend and replace)",

	Long: upsertLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'code'")
			cmd.Usage()
			os.Exit(1)
		}

		var code string

		if 0 < len(args) {

			code = args[0]

		}

		var globs []string

		if 1 < len(args) {

			globs = args[1:]

		}

		err = UpsertRun(code, globs)
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

	ohelp := UpsertCmd.HelpFunc()
	ousage := UpsertCmd.UsageFunc()
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

	UpsertCmd.SetHelpFunc(help)
	UpsertCmd.SetUsageFunc(usage)

}