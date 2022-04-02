package cmd

import "github.com/spf13/cobra"

func NewCommitteeListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all committees",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
