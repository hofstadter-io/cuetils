package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

var pipelineLong = `run file(s) through a pipeline of operations`

func init() {

	PipelineCmd.Flags().StringSliceVarP(&(flags.PipelineFlags.Tags), "tags", "t", nil, "tags to match for what to run")
}

func PipelineRun(globs []string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var PipelineCmd = &cobra.Command{

	Use: "pipeline <code> [files...]",

	Aliases: []string{
		"pipe",
		"dag",
	},

	Short: "run file(s) through a pipeline of operations",

	Long: pipelineLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		var globs []string

		if 0 < len(args) {

			globs = args[0:]

		}

		err = PipelineRun(globs)
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

	ohelp := PipelineCmd.HelpFunc()
	ousage := PipelineCmd.UsageFunc()
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

	PipelineCmd.SetHelpFunc(help)
	PipelineCmd.SetUsageFunc(usage)

}
