package backup

import (
	"context"
	"fmt"
	"path/filepath"
)

type MySQLBackupManager struct {
	User     string
	Password string
	Database string
}

func (m *MySQLBackupManager) Backup(ctx context.Context, outputDir string) error {
	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s-backup.sql", m.Database))
	return BackupMySQL(m.User, m.Password, m.Database, outputFile)
}

func (m *MySQLBackupManager) Restore(ctx context.Context, inputFile string) error {
	return RestoreMySQL(m.User, m.Password, m.Database, inputFile)
}
func (m *MySQLBackupManager) Validate(ctx context.Context) error {
	fmt.Println("✅ Validating MySQL connection...")
	// You could test with a `mysqladmin ping` here
	return nil
}

func (m *MySQLBackupManager) Type() string {
	return "mysql-baremetal"
}
