package cmd

import "github.com/spf13/cobra"

func NewCommitteeProposalCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new proposal",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
