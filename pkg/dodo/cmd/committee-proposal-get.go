package cmd

import "github.com/spf13/cobra"

func NewCommitteeProposalGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a proposal",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
