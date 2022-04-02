package cmd

import "github.com/spf13/cobra"

func NewCommitteeAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new committee",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
