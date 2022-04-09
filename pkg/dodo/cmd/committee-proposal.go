package cmd

import "github.com/spf13/cobra"

func NewCommitteeProposalCommand(*GlobalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "committee-proposal",
		Short: "Committee proposal commands",
	}

	cmd.AddCommand(NewCommitteeProposalDecryptApproveCommand())

	return cmd
}
