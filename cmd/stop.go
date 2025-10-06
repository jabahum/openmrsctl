package cmd

import (
	"fmt"

	"github.com/jabahum/openmrsctl/internal/services"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop OpenMRS and related services",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := services.GetManager()
		if err := manager.Stop(); err != nil {
			return fmt.Errorf("failed to stop OpenMRS services: %v", err)
		}
		fmt.Println("🛑 OpenMRS services stopped successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
