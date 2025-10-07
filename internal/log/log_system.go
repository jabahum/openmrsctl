package logs

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

type SystemdLogManager struct{}

func (s *SystemdLogManager) GetLogs(options LogOptions) (string, error) {
	serviceName, err := s.getServiceName(options.Component)
	if err != nil {
		return "", fmt.Errorf("failed to determine systemd service name for component '%s': %w", options.Component, err)
	}
	args := []string{"-u", serviceName}

	if options.TailLines > 0 {
		args = append(args, "-n", strconv.Itoa(options.TailLines))
	} else {
		args = append(args, "-n", "1000")
	}

	if !options.Since.IsZero() {
		args = append(args, "--since", options.Since.Format("2006-01-02 15:04:05"))
	}

	args = append(args, "--no-pager")

	cmd := exec.Command("journalctl", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("journalctl command failed for service '%s': %s (stderr: %s)", serviceName, err, stderr.String())
	}

	logs := stdout.String()

	if options.Level != "" {
		logs = filterLogsByLevel(logs, options.Level)
	}

	if options.GrepPattern != "" {
		logs = filterLogsByGrep(logs, options.GrepPattern)
	}

	return logs, nil
}

func (s *SystemdLogManager) FollowLogs(options LogOptions) error {
	// Implementation for following logs from systemd
	return nil
}

func (s *SystemdLogManager) SaveLogsToFile(filename string, options LogOptions) error {
	// Implementation for saving logs to a file from systemd
	return nil
}

func (s *SystemdLogManager) SetLogLevel(component string, level string, duration time.Duration) error {
	// Implementation for setting log level in systemd
	return nil
}

func (s *SystemdLogManager) BundleLogs(filename string, options LogOptions) (string, error) {
	// Implementation for bundling logs from systemd
	return "path/to/bundled/logs.tar.gz", nil
}
