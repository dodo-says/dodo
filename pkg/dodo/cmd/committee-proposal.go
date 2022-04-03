package cmd

import "github.com/spf13/cobra"

func NewCommitteeProposalCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal",
		Short: "Committee proposal commands",
	}

	cmd.AddCommand(NewCommitteeProposalCreateCommand())
	cmd.AddCommand(NewCommitteeProposalApproveCommand())
	cmd.AddCommand(NewCommitteeProposalGetCommand())
	cmd.AddCommand(NewCommitteeProposalListCommand())

	return cmd
}
