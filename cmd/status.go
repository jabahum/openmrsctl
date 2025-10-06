package cmd

import (
	"fmt"

	"github.com/jabahum/openmrsctl/internal/services"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show OpenMRS service status",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := services.GetManager()
		status, err := manager.Status()
		if err != nil {
			return fmt.Errorf("failed to check status: %v", err)
		}
		fmt.Println("📈 OpenMRS System Status:")
		fmt.Println(status)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
