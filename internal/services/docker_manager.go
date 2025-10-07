// internal/services/docker_manager.go
package services

import (
	"fmt"
	"os"
	"os/exec"
)

type DockerManager struct{}

// Start brings up the OpenMRS stack using docker compose.
func (d *DockerManager) Start() error {
	fmt.Println("🚀 Starting OpenMRS (Docker environment)...")
	cmd := exec.Command("bash", "-c", "docker compose up -d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Stop shuts down the OpenMRS stack.
func (d *DockerManager) Stop() error {
	fmt.Println("🛑 Stopping OpenMRS...")
	cmd := exec.Command("bash", "-c", "docker compose down")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Status shows running container statuses for OpenMRS-related services.
func (d *DockerManager) Status() (string, error) {
	cmd := exec.Command("bash", "-c", "docker compose ps --format 'table {{.Name}}\t{{.State}}\t{{.Status}}'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get status: %v", err)
	}
	return string(output), nil
}

// Restart restarts the OpenMRS containers.
func (d *DockerManager) Restart() error {
	fmt.Println("🔄 Restarting OpenMRS (Docker)...")
	if err := d.Stop(); err != nil {
		return fmt.Errorf("failed to stop containers: %v", err)
	}
	if err := d.Start(); err != nil {
		return fmt.Errorf("failed to start containers: %v", err)
	}
	return nil
}
