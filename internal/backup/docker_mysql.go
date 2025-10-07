package backup

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type DockerMySQLBackupManager struct {
	Container string
	User      string
	Password  string
	Database  string
}

func (m *DockerMySQLBackupManager) Backup(ctx context.Context, outputDir string) error {
	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s-backup.sql", m.Database))
	fmt.Printf("🐳 Backing up MySQL (Docker) container %s → %s\n", m.Container, outputFile)

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cmd := exec.CommandContext(ctx, "docker", "exec", m.Container,
		"mysqldump", "-u", m.User, fmt.Sprintf("-p%s", m.Password), m.Database)
	cmd.Stdout = outFile
	return cmd.Run()
}

func (m *DockerMySQLBackupManager) Restore(ctx context.Context, inputFile string) error {
	fmt.Printf("🐳 Restoring MySQL in container %s from %s\n", m.Container, inputFile)
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	cmd := exec.CommandContext(ctx, "docker", "exec", "-i", m.Container,
		"mysql", "-u", m.User, fmt.Sprintf("-p%s", m.Password), m.Database)
	cmd.Stdin = inFile
	return cmd.Run()
}

func (m *DockerMySQLBackupManager) Validate(ctx context.Context) error {
	fmt.Printf("🔍 Checking Docker container: %s\n", m.Container)
	return exec.CommandContext(ctx, "docker", "inspect", m.Container).Run()
}

func (m *DockerMySQLBackupManager) Type() string {
	return "mysql-docker"
}
