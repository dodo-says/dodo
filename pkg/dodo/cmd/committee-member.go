package cmd

import "github.com/spf13/cobra"

func NewCommitteeMemberCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "committee-member",
		Short: "Committee member commands",
	}

	cmd.AddCommand(NewCommitteeMemberAddCommand())
	cmd.AddCommand(NewCommitteeMemberRemoveCommand())
	cmd.AddCommand(NewCommitteeMemberListCommand())

	return cmd
}
