package cmd

import "github.com/spf13/cobra"

func NewCommitteeMemberRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a member from a committee",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
