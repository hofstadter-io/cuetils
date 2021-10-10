package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var replaceLong = `apply the replace from the original to the glob file(s)`

func ReplaceRun(orig string, globs []string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var ReplaceCmd = &cobra.Command{

	Use: "replace <orig> <glob>",

	Short: "apply the replace from the original to the glob file(s)",

	Long: replaceLong,

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

		var globs []string

		if 1 < len(args) {

			globs = args[1:]

		}

		err = ReplaceRun(orig, globs)
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

	ohelp := ReplaceCmd.HelpFunc()
	ousage := ReplaceCmd.UsageFunc()
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

	ReplaceCmd.SetHelpFunc(help)
	ReplaceCmd.SetUsageFunc(usage)

}
