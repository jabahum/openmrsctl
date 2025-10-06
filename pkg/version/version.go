package version

import (
	"fmt"
	"runtime"
)

// These variables are set at build time using -ldflags
var (
	Version   = "dev"
	GitCommit = "none"
	BuildDate = "unknown"
)

// Info returns formatted version info
func Info() string {
	return fmt.Sprintf(`openmrsctl:
  Version:    %s
  Git Commit: %s
  Build Date: %s
  Go Version: %s
  OS/Arch:    %s/%s
`, Version, GitCommit, BuildDate, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
