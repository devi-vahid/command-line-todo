package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var dataFile string

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A simple CLI task manager",
	Long:  "Manage tasks in a CSV data file using commands like add, list, complete, and delete.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&dataFile,
		"file",
		"f",
		"tasks.csv",
		"Path to the CSV file for task storage",
	)
}
