package cmd

import "github.com/spf13/cobra"

func NewCommitteeCommand(globalOptions *GlobalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "committee",
		Short: "Operations about committee",
	}

	// committee add
	cmd.AddCommand(NewCommitteeAddCommand(globalOptions))
	// committee get
	cmd.AddCommand(NewCommitteeGetCommand(globalOptions))
	// committee list
	cmd.AddCommand(NewCommitteeListCommand(globalOptions))
	// committee remove
	cmd.AddCommand(NewCommitteeRemoveCommand(globalOptions))
	// committee member
	cmd.AddCommand(NewCommitteeMemberCommand(globalOptions))
	// committee proposal
	cmd.AddCommand(NewCommitteeProposalCommand(globalOptions))

	return cmd
}
