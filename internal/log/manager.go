package logs

import (
	"os/exec"
	"time"
)

// lets create an interface for logs for both docker and systemd

type LogOptions struct {
	Component   string
	TailLines   int
	Since       time.Time
	Level       string
	GrepPattern string
}
type LogManager interface {
	GetLogs(options LogOptions) (string, error)
	FollowLogs(options LogOptions) error
	SaveLogsToFile(filename string, options LogOptions) error
	SetLogLevel(component string, level string, duration time.Duration) error
	BundleLogs(filename string, options LogOptions) (string, error)
}

func GetLogManager() LogManager {
	if isDockerAvailable() {
		return &DockerLogManager{}
	}
	return &SystemdLogManager{}
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
