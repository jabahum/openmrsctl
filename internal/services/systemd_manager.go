// internal/services/systemd_manager.go
package services

import (
	"fmt"
	"os/exec"
)

type SystemdManager struct{}

func (s *SystemdManager) Start() error {
	fmt.Println("Starting OpenMRS services via systemd...")
	if err := exec.Command("sudo", "systemctl", "start", "tomcat9").Run(); err != nil {
		return fmt.Errorf("failed to start tomcat9: %v", err)
	}
	if err := exec.Command("sudo", "systemctl", "start", "mysql").Run(); err != nil {
		return fmt.Errorf("failed to start mysql: %v", err)
	}
	return nil
}

func (s *SystemdManager) Stop() error {
	fmt.Println("Stopping OpenMRS services via systemd...")
	exec.Command("sudo", "systemctl", "stop", "tomcat9").Run()
	exec.Command("sudo", "systemctl", "stop", "mysql").Run()
	return nil
}

func (s *SystemdManager) Restart() error {
	fmt.Println("Restarting OpenMRS services via systemd...")
	exec.Command("sudo", "systemctl", "restart", "tomcat9").Run()
	exec.Command("sudo", "systemctl", "restart", "mysql").Run()
	return nil
}

func (s *SystemdManager) Status() (string, error) {
	out, err := exec.Command("systemctl", "is-active", "tomcat9").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("status check failed: %v", err)
	}
	return fmt.Sprintf("Tomcat9: %s", string(out)), nil
}

func (s *SystemdManager) Logs() (string, error) {
	out, err := exec.Command("sudo", "journalctl", "-u", "tomcat9", "--no-pager", "-n", "20").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get logs: %v", err)
	}
	return string(out), nil
}
