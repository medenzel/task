package cmd

import (
	"fmt"
	"os"

	"github.com/medenzel/task/models"
	"github.com/spf13/cobra"
)

type CommandService struct {
	TS      models.TaskService
	RootCmd *cobra.Command
}

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
}

func (cs *CommandService) SetRoot() {
	cs.RootCmd = rootCmd
}

func (cs *CommandService) Execute() {
	cs.SetRoot()
	cs.SetCommands()
	if err := cs.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
