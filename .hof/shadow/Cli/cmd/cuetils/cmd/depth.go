package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var depthLong = `calculate the depth of file(s)`

func DepthRun(globs []string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var DepthCmd = &cobra.Command{

	Use: "depth [files...]",

	Short: "calculate the depth of file(s)",

	Long: depthLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		var globs []string

		if 0 < len(args) {

			globs = args[0:]

		}

		err = DepthRun(globs)
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

	ohelp := DepthCmd.HelpFunc()
	ousage := DepthCmd.UsageFunc()
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

	DepthCmd.SetHelpFunc(help)
	DepthCmd.SetUsageFunc(usage)

}
