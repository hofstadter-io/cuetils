package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
	"github.com/hofstadter-io/cuetils/structural"
)

var validateLong = `validate file(s) with schema`

func ValidateRun(schema string, globs []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	results, err := structural.ValidateGlobs(schema, globs, flags.RootPflags)
	if err != nil {
		return err
	}

	for _, r := range results {
		fmt.Printf("%s\n----------------------\n%s\n\n", r.Filename, r.Content)
	}

	if len(results) > 0 {
		err = fmt.Errorf("Errors in %d file(s)", len(results))
	}

	return err
}

var ValidateCmd = &cobra.Command{

	Use: "validate <schema> [files...]",

	Aliases: []string{
		"v",
	},

	Short: "validate file(s) with schema",

	Long: validateLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'schema'")
			cmd.Usage()
			os.Exit(1)
		}

		var schema string

		if 0 < len(args) {

			schema = args[0]

		}

		var globs []string

		if 1 < len(args) {

			globs = args[1:]

		}

		err = ValidateRun(schema, globs)
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