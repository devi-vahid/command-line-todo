package cmd

import (
	"fmt"
	"time"

	"github.com/devi-vahid/command-line-todo/internal/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		description := args[0]
		tasks, err := task.LoadTasks(dataFile)

		if nil != err {
			return err
		}

		newTask := task.Task{
			ID:          len(tasks) + 1,
			Description: description,
			CreatedAt:   time.Now(),
			IsComplete:  false,
		}

		tasks = append(tasks, newTask)

		err = task.SaveTasks(dataFile, tasks)

		if nil != err {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Added task: %s\n", description)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
