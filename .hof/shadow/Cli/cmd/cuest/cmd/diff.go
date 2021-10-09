package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var diffLong = `calculate the diff from the original to the glob file(s)`

func DiffRun(orig string, glob string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var DiffCmd = &cobra.Command{

	Use: "diff <orig> <glob>",

	Short: "calculate the diff from the original to the glob file(s)",

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

		var glob string

		if 1 < len(args) {

			glob = args[1]

		}

		err = DiffRun(orig, glob)
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