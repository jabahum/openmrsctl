// cmd/init.go
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jabahum/openmrsctl/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize OpenMRS environment for openmrsctl",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("⚙️ Initializing OpenMRS environment...")

		reader := bufio.NewReader(os.Stdin)

		// 1️⃣ Deployment Type
		fmt.Print("Select deployment type (docker/baremetal) [baremetal]: ")
		deployment, _ := reader.ReadString('\n')
		deployment = strings.TrimSpace(deployment)
		if deployment == "" {
			deployment = "baremetal"
		}

		// 2️⃣ OpenMRS Home Directory
		fmt.Print("Enter OpenMRS home directory [/var/lib/tomcat9/webapps/openmrs]: ")
		openmrsHome, _ := reader.ReadString('\n')
		openmrsHome = strings.TrimSpace(openmrsHome)
		if openmrsHome == "" {
			openmrsHome = "/var/lib/tomcat9/webapps/openmrs"
		}

		// 3️⃣ Tomcat Service Name
		fmt.Print("Enter Tomcat service name [tomcat9]: ")
		tomcatService, _ := reader.ReadString('\n')
		tomcatService = strings.TrimSpace(tomcatService)
		if tomcatService == "" {
			tomcatService = "tomcat9"
		}

		// 4️⃣ MySQL Service Name
		fmt.Print("Enter MySQL service name [mysql]: ")
		mysqlService, _ := reader.ReadString('\n')
		mysqlService = strings.TrimSpace(mysqlService)
		if mysqlService == "" {
			mysqlService = "mysql"
		}

		// 5️⃣ Backup Directory
		fmt.Print("Enter backup directory [./backups]: ")
		backupDir, _ := reader.ReadString('\n')
		backupDir = strings.TrimSpace(backupDir)
		if backupDir == "" {
			backupDir = "./backups"
		}

		// 6️⃣ Logs Directory
		fmt.Print("Enter logs directory [./logs]: ")
		logsDir, _ := reader.ReadString('\n')
		logsDir = strings.TrimSpace(logsDir)
		if logsDir == "" {
			logsDir = "./logs"
		}

		// Create folders
		dirs := []string{backupDir, logsDir, "./tmp"}
		for _, d := range dirs {
			if err := os.MkdirAll(d, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %v", d, err)
			}
		}

		// Save config
		conf := config.Config{
			Deployment:    deployment,
			OpenMRSHome:   openmrsHome,
			TomcatService: tomcatService,
			MySQLService:  mysqlService,
			BackupDir:     backupDir,
			LogsDir:       logsDir,
		}
		if err := config.SaveConfig(conf); err != nil {
			return fmt.Errorf("failed to save config: %v", err)
		}

		fmt.Println("✅ Initialization completed successfully!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
