package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/cuetils/cmd/cuest/verinfo"
)

const versionMessage = `
Version:     v%s
Commit:      %s

BuildDate:   %s
GoVersion:   %s
OS / Arch:   %s %s


Author:   Hofstadter, Inc
Homepage: https://docs.hofstadter.io
GitHub:   https://github.com/hofstadter-io/cuetils

`

var VersionLong = `Print the build version for cuest`

var VersionCmd = &cobra.Command{

	Use: "version",

	Aliases: []string{
		"ver",
	},

	Short: "print the version",

	Long: VersionLong,

	Run: func(cmd *cobra.Command, args []string) {

		s, e := os.UserConfigDir()
		fmt.Printf("cuest ConfigDir %q %v\n", filepath.Join(s, "cuest"), e)

		fmt.Printf(
			versionMessage,
			verinfo.Version,
			verinfo.Commit,
			verinfo.BuildDate,
			verinfo.GoVersion,
			verinfo.BuildOS,
			verinfo.BuildArch,
		)
	},
}

func init() {
	help := VersionCmd.HelpFunc()
	usage := VersionCmd.UsageFunc()

	VersionCmd.SetHelpFunc(help)
	VersionCmd.SetUsageFunc(usage)

}
