package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var pickLong = `pick the original from the glob file(s)`

func PickRun(pick string, globs []string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var PickCmd = &cobra.Command{

	Use: "pick <pick> <glob>",

	Short: "pick the original from the glob file(s)",

	Long: pickLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'pick'")
			cmd.Usage()
			os.Exit(1)
		}

		var pick string

		if 0 < len(args) {

			pick = args[0]

		}

		var globs []string

		if 1 < len(args) {

			globs = args[1:]

		}

		err = PickRun(pick, globs)
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

	ohelp := PickCmd.HelpFunc()
	ousage := PickCmd.UsageFunc()
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

	PickCmd.SetHelpFunc(help)
	PickCmd.SetUsageFunc(usage)

}
