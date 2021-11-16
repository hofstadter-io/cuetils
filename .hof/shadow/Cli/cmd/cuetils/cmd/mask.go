package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var maskLong = `mask from file(s) with code`

func MaskRun(code string, globs []string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var MaskCmd = &cobra.Command{

	Use: "mask <code> [files...]",

	Short: "mask from file(s) with code",

	Long: maskLong,

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

		err = MaskRun(code, globs)
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

	ohelp := MaskCmd.HelpFunc()
	ousage := MaskCmd.UsageFunc()
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

	MaskCmd.SetHelpFunc(help)
	MaskCmd.SetUsageFunc(usage)

}
