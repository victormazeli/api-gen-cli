package main

import (
	"github.com/spf13/cobra"
)

var noInput bool

// NewRootCmd Root command function.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "api-gen",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.PersistentFlags().BoolVar(&noInput, "no-input", false, "Disable interactivity")
	addSubCommands(rootCmd)
	return rootCmd
}

// Function to add sub commands.
func addSubCommands(cmd *cobra.Command) {
	subcmds := []*cobra.Command{
		generateCmd(),
	}
	for _, subcmd := range subcmds {
		cmd.AddCommand(subcmd)
	}
}
