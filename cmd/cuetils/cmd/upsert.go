package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/cuetils/pkg/structural"
)

var upsertLong = `apply the upsert from the original to the glob file(s)`

func UpsertRun(orig string, globs []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	upserts, err := structural.Upsert(orig, globs)
	if err != nil {
		return err
	}

	for _, u := range upserts {
		fmt.Printf("%s\n----------------------\n%s\n\n", u.Filename, u.Content)
	}

	return err
}

var UpsertCmd = &cobra.Command{

	Use: "upsert <orig> <glob>",

	Short: "apply the upsert from the original to the glob file(s)",

	Long: upsertLong,

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

		err = UpsertRun(orig, globs)
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
