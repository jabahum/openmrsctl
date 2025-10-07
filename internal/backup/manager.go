package backup

import "context"

// BackupManager provides a uniform interface for backup operations.
type BackupManager interface {
	Backup(ctx context.Context, outputDir string) error
	Restore(ctx context.Context, inputFile string) error
	Validate(ctx context.Context) error
	Type() string
}
