package cmd

import "github.com/spf13/cobra"

func NewCommitteeMemberListCommand(*GlobalOptions) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List all committee members",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
