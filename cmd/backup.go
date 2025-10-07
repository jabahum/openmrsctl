package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jabahum/openmrsctl/internal/backup"
	"github.com/spf13/cobra"
)

var (
	outputDir  string
	dockerMode bool
	container  string
	user       string
	password   string
	database   string
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup OpenMRS databases and related data",
	Long:  `Create a backup of the OpenMRS database and related files for bare metal or Docker-based installations.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		timestamp := time.Now().Format("20060102-150405")
		output := filepath.Join(outputDir, timestamp)
		os.MkdirAll(output, 0755)

		var manager backup.BackupManager
		if dockerMode {
			manager = &backup.DockerMySQLBackupManager{
				Container: container,
				User:      user,
				Password:  password,
				Database:  database,
			}
		} else {
			manager = &backup.MySQLBackupManager{
				User:     user,
				Password: password,
				Database: database,
			}
		}

		fmt.Printf("Starting backup using %s...\n", manager.Type())

		if err := manager.Validate(ctx); err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}

		if err := manager.Backup(ctx, output); err != nil {
			return fmt.Errorf("backup failed: %w", err)
		}

		fmt.Printf("✅ Backup completed successfully → %s\n", output)
		return nil
	},
}

func init() {
	backupCmd.Flags().StringVarP(&outputDir, "output", "o", "./backups", "Output directory for backups")
	backupCmd.Flags().BoolVar(&dockerMode, "docker", false, "Use Docker container mode")
	backupCmd.Flags().StringVar(&container, "container", "mysql", "Docker container name for MySQL")
	backupCmd.Flags().StringVar(&user, "user", "root", "Database username")
	backupCmd.Flags().StringVar(&password, "password", "password", "Database password")
	backupCmd.Flags().StringVar(&database, "db", "openmrs", "Database name")
	rootCmd.AddCommand(backupCmd)
}
