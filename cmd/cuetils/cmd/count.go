package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
	"github.com/hofstadter-io/cuetils/structural"
)

var countLong = `count nodes in file(s)`

func CountRun(globs []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	counts, err := structural.CountGlobs(globs, &flags.RootPflags)
	if err != nil {
		return err
	}

	for _, c := range counts {
		fmt.Println(c.Filename, c.Count)
	}

	return err
}

var CountCmd = &cobra.Command{

	Use: "count [files...]",

	Short: "count nodes in file(s)",

	Long: countLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		var globs []string

		if 0 < len(args) {

			globs = args[0:]

		}

		err = CountRun(globs)
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

	ohelp := CountCmd.HelpFunc()
	ousage := CountCmd.UsageFunc()
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

	CountCmd.SetHelpFunc(help)
	CountCmd.SetUsageFunc(usage)

}
