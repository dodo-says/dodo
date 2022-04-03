package cmd

import "github.com/spf13/cobra"

func NewCommitteeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "committee",
		Short: "Operations about committee",
	}
	cmd.AddCommand(NewCommitteeAddCommand())
	cmd.AddCommand(NewCommitteeListCommand())
	cmd.AddCommand(NewCommitteeRemoveCommand())
	cmd.AddCommand(NewCommitteeMemberCommand())
	cmd.AddCommand(NewCommitteeProposalCommand())
	return cmd
}
