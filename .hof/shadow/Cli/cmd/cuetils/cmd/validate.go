package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var validateLong = `validate with the original against the glob file(s)`

func ValidateRun(orig string, globs []string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var ValidateCmd = &cobra.Command{

	Use: "validate <orig> <glob>",

	Short: "validate with the original against the glob file(s)",

	Long: validateLong,

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

		err = ValidateRun(orig, globs)
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

	ohelp := ValidateCmd.HelpFunc()
	ousage := ValidateCmd.UsageFunc()
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

	ValidateCmd.SetHelpFunc(help)
	ValidateCmd.SetUsageFunc(usage)

}
