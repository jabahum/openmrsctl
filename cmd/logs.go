package cmd

import (
	"fmt"

	"github.com/jabahum/openmrsctl/internal/services"
	"github.com/spf13/cobra"
)

var tailLines int

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "View recent OpenMRS logs",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := services.GetManager()
		logs, err := manager.Logs()
		if err != nil {
			return fmt.Errorf("failed to fetch logs: %v", err)
		}
		fmt.Println("🧩 OpenMRS Logs:")
		fmt.Println(logs)
		return nil
	},
}

func init() {
	logsCmd.Flags().IntVarP(&tailLines, "lines", "n", 20, "Number of log lines to display")
	rootCmd.AddCommand(logsCmd)
}
