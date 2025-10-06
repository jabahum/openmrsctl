package cmd

import (
	"fmt"

	"github.com/jabahum/openmrsctl/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(version.Info())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
