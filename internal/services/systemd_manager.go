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

func (s *SystemdManager) Backup(file string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf(`
		mkdir -p $(dirname %s) &&
		mysqldump -u root -p openmrs > /tmp/openmrs.sql &&
		tar -czf %s /tmp/openmrs.sql /var/lib/tomcat9/webapps/openmrs &&
		rm /tmp/openmrs.sql
	`, file))
	return cmd.Run()
}

func (s *SystemdManager) Restore(file string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf(`
		tar -xzf %s -C /tmp &&
		mysql -u root -p openmrs < /tmp/openmrs.sql &&
		cp -r /tmp/openmrs /var/lib/tomcat9/webapps/
	`, file))
	return cmd.Run()
}
