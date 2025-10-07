package cmd

import (
	"fmt"
	"time"

	logs "github.com/jabahum/openmrsctl/internal/log"
	"github.com/spf13/cobra"
)

var logFollow bool
var logTail int
var logComponent string
var logLevel string
var logSince string

var logCmd = &cobra.Command{
	Use:   "log [options]",
	Short: "Stream or display OpenMRS logs, supporting filtering and multiple components.",
	Long: `Retrieve, filter, and stream logs from the running OpenMRS instance.
    Automatically detects Docker or Bare-Metal environments.
    
    Example:
    openmrsctl log -f --level ERROR --component db
    openmrsctl log -n 50 --since 1h
    `,
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := logs.GetLogManager()

		// --- 1. Construct LogOptions from flags ---
		options := logs.LogOptions{
			TailLines: logTail,
			Component: logComponent,
			Level:     logLevel,
			// GrepPattern will be an argument or a separate flag if needed
		}

		// Parse --since flag
		if logSince != "" {
			duration, err := time.ParseDuration(logSince)
			if err != nil {
				return fmt.Errorf("invalid duration for --since flag: %w", err)
			}
			// Calculate the time point: Current Time - Duration
			options.Since = time.Now().Add(-duration)
		}

		// --- 2. Execute based on the --follow flag ---
		if logFollow {
			fmt.Println("🔎 Streaming OpenMRS logs. Press Ctrl+C to stop...")
			return manager.FollowLogs(options)
			// Note: FollowLogs implementation pipes output directly to stdout.
		}

		// Historical Log Retrieval
		logs, err := manager.GetLogs(options)
		if err != nil {
			return fmt.Errorf("failed to fetch historical logs: %w", err)
		}

		fmt.Println("🧩 OpenMRS Logs:")
		fmt.Println(logs)
		return nil
	},
}

func init() {
	// Flag for number of lines (replaces the old tailLines variable)
	logCmd.Flags().IntVarP(&logTail, "lines", "n", 20, "Number of log lines to display (for non-stream mode)")

	// Flag for real-time streaming
	logCmd.Flags().BoolVarP(&logFollow, "follow", "f", false, "Stream the logs in real-time (like tail -f)")

	// Flag for component selection
	logCmd.Flags().StringVarP(&logComponent, "component", "c", "app", "Specify the component/service (e.g., app, db, tomcat)")

	// Flag for log level filtering
	logCmd.Flags().StringVar(&logLevel, "level", "", "Filter logs by minimum level (e.g., DEBUG, INFO, WARN, ERROR)")

	// Flag for time-based filtering
	logCmd.Flags().StringVar(&logSince, "since", "", "Show logs since a duration (e.g., 5m, 1h, 2d)")

	rootCmd.AddCommand(logCmd)
}
