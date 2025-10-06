package cmd

import (
	"fmt"
	"time"

	"github.com/jabahum/openmrsctl/internal/services"
	"github.com/spf13/cobra"
)

var backupDir string

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup OpenMRS database and application data",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := services.GetManager()
		timestamp := time.Now().Format("20060102-150405")
		file := fmt.Sprintf("%s/openmrs-backup-%s.tar.gz", backupDir, timestamp)

		fmt.Printf("🗄️  Starting backup: %s\n", file)
		if err := manager.Backup(file); err != nil {
			return fmt.Errorf("backup failed: %v", err)
		}

		fmt.Println("✅ Backup completed successfully!")
		return nil
	},
}

func init() {
	backupCmd.Flags().StringVarP(&backupDir, "output", "o", "./backups", "Backup output directory")
	rootCmd.AddCommand(backupCmd)
}
