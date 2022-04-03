package cmd

import "github.com/spf13/cobra"

func NewCommitteeProposalListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all committee proposals",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
