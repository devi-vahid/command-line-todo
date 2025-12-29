package cmd

import (
	"fmt"
	"strconv"

	"github.com/devi-vahid/command-line-todo/internal/task"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "mark task as completed",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskId, err := strconv.Atoi(args[0])

		if nil != err {
			return err
		}

		tasks, err := task.LoadTasks(dataFile)

		if nil != err {
			return err
		}

		for i := 0; i < len(tasks); i++ {
			if taskId == tasks[i].ID {
				tasks[i].IsComplete = true
			}
		}

		err = task.SaveTasks(dataFile, tasks)

		if nil != err {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Marked task as complete: %d\n", taskId)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
