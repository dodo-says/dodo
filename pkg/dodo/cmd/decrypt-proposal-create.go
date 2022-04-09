package cmd

import "github.com/spf13/cobra"

func NewDecryptProposalCreateCommand(*GlobalOptions) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new decrypt proposal",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}
