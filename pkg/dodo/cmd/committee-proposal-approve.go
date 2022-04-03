package cmd

import (
	"github.com/spf13/cobra"
)

func NewCommitteeProposalApproveCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "approve",
		Short: "Approve a proposal",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}
