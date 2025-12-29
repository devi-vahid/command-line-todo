package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"

	"github.com/devi-vahid/command-line-todo/internal/task"
)

var showAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := task.LoadTasks(dataFile)
		if nil != err {
			return err
		}

		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)

		if showAll {
			fmt.Fprintln(w, "ID\tTask\tCreatedAt\tIsComplete")
		} else {
			fmt.Fprintln(w, "ID\tTask\tCreatedAt")
		}

		for _, task := range tasks {
			if !showAll && task.IsComplete {
				continue
			}
			if showAll {
				fmt.Fprintf(w, "%d\t%s\t%s\t%v\n",
					task.ID,
					task.Description,
					timediff.TimeDiff(task.CreatedAt),
					task.IsComplete,
				)
			} else {
				fmt.Fprintf(w, "%d\t%s\t%s\n",
					task.ID,
					task.Description,
					timediff.TimeDiff(task.CreatedAt),
				)
			}
		}
		w.Flush()
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "show completed tasks too")
	rootCmd.AddCommand(listCmd)
}
