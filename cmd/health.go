package cmd

import (
	"fmt"

	"github.com/jabahum/openmrsctl/internal/health"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Perform health checks on OpenMRS services",
	RunE: func(cmd *cobra.Command, args []string) error {
		results := health.RunChecks()
		fmt.Println("🔍 OpenMRS Health Check Results:")
		for _, r := range results {
			fmt.Printf(" - %s: %s\n", r.Name, r.Status)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}
