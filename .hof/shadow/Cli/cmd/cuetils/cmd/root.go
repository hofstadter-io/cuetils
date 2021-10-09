package cmd

import (
	"fmt"
	"os"

	"log"
	"runtime/pprof"

	"github.com/hofstadter-io/hof/script/runtime"
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/cuetils/cmd/cuetils/flags"
)

var cuetilsLong = `CUE Utilites for bulk ETL, diff, query, and other operations on data and config`

func init() {

	RootCmd.PersistentFlags().IntVarP(&(flags.RootPflags.Maxiter), "maxiter", "m", 0, "maximum iterations for recursion")
	RootCmd.PersistentFlags().StringVarP(&(flags.RootPflags.Out), "out", "", "cue", "output encoding [cue,yaml,json]")
}

func RootPersistentPostRun(args []string) (err error) {

	WaitPrintUpdateAvailable()

	return err
}

var RootCmd = &cobra.Command{

	Use: "cuetils",

	Short: "CUE Utilites for bulk ETL, diff, query, and other operations on data and config",

	Long: cuetilsLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = RootPersistentPostRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func RootInit() {
	extra := func(cmd *cobra.Command) bool {

		return false
	}

	ohelp := RootCmd.HelpFunc()
	ousage := RootCmd.UsageFunc()
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

	RootCmd.SetHelpFunc(help)
	RootCmd.SetUsageFunc(usage)

	RootCmd.AddCommand(UpdateCmd)

	RootCmd.AddCommand(VersionCmd)

	RootCmd.AddCommand(CompletionCmd)

	RootCmd.AddCommand(DepthCmd)
	RootCmd.AddCommand(DiffCmd)
	RootCmd.AddCommand(PatchCmd)
	RootCmd.AddCommand(PickCmd)
	RootCmd.AddCommand(MaskCmd)

}

func RunExit() {
	if err := RunErr(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RunInt() int {
	if err := RunErr(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func RunErr() error {

	if fn := os.Getenv("CUETILS_CPU_PROFILE"); fn != "" {
		f, err := os.Create(fn)
		if err != nil {
			log.Fatal("Could not create file for CPU profile:", err)
		}
		defer f.Close()

		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal("Could not start CPU profile process:", err)
		}

		defer pprof.StopCPUProfile()
	}

	RootInit()
	return RootCmd.Execute()
}

func CallTS(ts *runtime.Script, args []string) error {
	RootCmd.SetArgs(args)

	err := RootCmd.Execute()
	ts.Check(err)

	return err
}
