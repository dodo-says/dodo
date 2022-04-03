package cmd

import (
	"github.com/spf13/cobra"
)

func NewCommitteeRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a member from a committee",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
