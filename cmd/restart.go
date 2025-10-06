package cmd

import (
	"fmt"

	"github.com/jabahum/openmrsctl/internal/services"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart OpenMRS services",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := services.GetManager()
		fmt.Println("🔄 Restarting OpenMRS...")
		if err := manager.Restart(); err != nil {
			return fmt.Errorf("restart failed: %v", err)
		}
		fmt.Println("✅ OpenMRS restarted successfully!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}
