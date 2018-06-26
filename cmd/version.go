package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("github-release-tool v" + version)
	},
}
