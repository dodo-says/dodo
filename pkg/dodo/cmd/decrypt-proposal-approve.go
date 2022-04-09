package cmd

import (
	"github.com/spf13/cobra"
)

func NewDecryptProposalApproveCommand(*GlobalOptions) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "approve",
		Short: "Approve a decrypt proposal",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}
