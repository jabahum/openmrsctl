package cmd

import (
	"fmt"

	"github.com/jabahum/openmrsctl/internal/services"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore <backup-file>",
	Short: "Restore OpenMRS from a backup file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file := args[0]
		manager := services.GetManager()

		fmt.Printf("♻️  Restoring from backup: %s\n", file)
		if err := manager.Restore(file); err != nil {
			return fmt.Errorf("restore failed: %v", err)
		}

		fmt.Println("✅ Restore completed successfully!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
