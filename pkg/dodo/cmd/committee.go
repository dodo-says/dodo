package cmd

import "github.com/spf13/cobra"

func NewCommitteeCommand(globalOptions *GlobalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "committee",
		Short: "Operations about committee",
	}
	cmd.AddCommand(NewCommitteeAddCommand(globalOptions))
	cmd.AddCommand(NewCommitteeListCommand(globalOptions))
	cmd.AddCommand(NewCommitteeRemoveCommand(globalOptions))
	cmd.AddCommand(NewCommitteeMemberCommand(globalOptions))
	cmd.AddCommand(NewCommitteeProposalCommand(globalOptions))
	return cmd
}
