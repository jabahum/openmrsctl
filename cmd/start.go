package cmd

import (
	"fmt"

	"github.com/jabahum/openmrsctl/internal/services"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start OpenMRS and related services",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := services.GetManager()
		if err := manager.Start(); err != nil {
			return fmt.Errorf("failed to start OpenMRS: %v", err)
		}
		fmt.Println("✅ OpenMRS services started successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
