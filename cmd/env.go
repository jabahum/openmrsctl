package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Show environment and OpenMRS info",
	RunE: func(cmd *cobra.Command, args []string) error {
		// OS
		fmt.Println("🌍 Environment Info:")
		out, _ := exec.Command("uname", "-a").Output()
		fmt.Printf("OS: %s\n", out)

		// Java
		out, _ = exec.Command("java", "-version").CombinedOutput()
		fmt.Printf("Java: %s\n", out)

		// MySQL
		out, _ = exec.Command("mysql", "--version").Output()
		fmt.Printf("MySQL: %s\n", out)

		// OpenMRS Version
		resp, err := exec.Command("curl", "-s", "http://localhost:8080/openmrs/ws/rest/v1/system/version").Output()
		if err == nil {
			fmt.Printf("OpenMRS Version: %s\n", resp)
		} else {
			fmt.Println("OpenMRS Version: Unknown")
		}

		// Deployment type
		deployment := "Bare Metal"
		if isDockerAvailable() {
			deployment = "Docker"
		}
		fmt.Printf("Deployment Type: %s\n", deployment)

		return nil
	},
}

func isDockerAvailable() bool {
	_, err := exec.LookPath("docker")
	return err == nil
}

func init() {
	rootCmd.AddCommand(envCmd)
}
