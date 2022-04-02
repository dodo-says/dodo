package cmd

import "github.com/spf13/cobra"

func NewCommitteeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "committee",
		Short: "Operations about committee",
	}
	cmd.AddCommand(NewCommitteeAddCommand())
	cmd.AddCommand(NewCommitteeListCommand())
	return cmd
}
