// internal/services/manager.go
package services

import "os/exec"

type ServiceManager interface {
	Start() error
	Stop() error
	Restart() error
	Status() (string, error)
	Logs() (string, error)
	Backup(file string) error
	Restore(file string) error
}

func GetManager() ServiceManager {
	if isDockerAvailable() {
		return &DockerManager{}
	}
	return &SystemdManager{}
}

func isDockerAvailable() bool {
	_, err := exec.LookPath("docker")
	if err != nil {
		return false
	}
	cmd := exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
