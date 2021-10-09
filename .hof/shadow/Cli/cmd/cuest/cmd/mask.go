package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var maskLong = `mask the original from the glob file(s)`

func MaskRun(mask string, glob string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var MaskCmd = &cobra.Command{

	Use: "mask <mask> <glob>",

	Short: "mask the original from the glob file(s)",

	Long: maskLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'mask'")
			cmd.Usage()
			os.Exit(1)
		}

		var mask string

		if 0 < len(args) {

			mask = args[0]

		}

		var glob string

		if 1 < len(args) {

			glob = args[1]

		}

		err = MaskRun(mask, glob)
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
