package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Args:  cobra.MinimumNArgs(1),
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Args:  cobra.ExactArgs(1),
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Args:  cobra.NoArgs,
}

func (cs *CommandService) SetCommands() {
	addCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return cs.TS.AddTask(strings.Join(args, " "))
	}
	doCmd.RunE = func(cmd *cobra.Command, args []string) error {
		ID, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("do command: %w", err)
		}
		return cs.TS.DoTask(ID)
	}
	listCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return cs.TS.ListTasks()
	}

	cs.RootCmd.AddCommand(addCmd)
	cs.RootCmd.AddCommand(doCmd)
	cs.RootCmd.AddCommand(listCmd)
}
