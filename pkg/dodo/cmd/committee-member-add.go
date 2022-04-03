package cmd

import "github.com/spf13/cobra"

func NewCommitteeMemberAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new member to the committee",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
