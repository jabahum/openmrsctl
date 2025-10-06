// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openmrsctl",
	Short: "A lightweight DevOps and maintenance tool for OpenMRS environments",
	Long:  `openmrsctl helps developers and system admins manage, monitor, and automate OpenMRS deployments — on bare metal or Docker.`,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
