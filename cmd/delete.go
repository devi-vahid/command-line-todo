package cmd

import (
	"github.com/devi-vahid/go-todo-list/internal/task"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "delete a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		tasks, err := task.LoadTasks(dataFile)
		if nil != err {
			return err
		}
		for _, task := range tasks {

		}

		return nil
	},
}
