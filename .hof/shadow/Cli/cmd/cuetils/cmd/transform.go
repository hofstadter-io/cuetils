package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var transformLong = `transform file(s) with code`

func TransformRun(code string, globs []string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var TransformCmd = &cobra.Command{

	Use: "transform <code> [files...]",

	Short: "transform file(s) with code",

	Long: transformLong,

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

		err = TransformRun(code, globs)
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

	ohelp := TransformCmd.HelpFunc()
	ousage := TransformCmd.UsageFunc()
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

	TransformCmd.SetHelpFunc(help)
	TransformCmd.SetUsageFunc(usage)

}
