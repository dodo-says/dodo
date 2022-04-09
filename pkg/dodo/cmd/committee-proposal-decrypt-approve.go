package cmd

import (
	"github.com/spf13/cobra"
)

func NewCommitteeProposalDecryptApproveCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "decrypt-approve",
		Short: "Approve a proposal",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}
