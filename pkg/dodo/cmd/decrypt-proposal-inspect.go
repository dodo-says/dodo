package cmd

import "github.com/spf13/cobra"

func NewDecryptProposalInspectCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "inspect",
		Short: "inspect the detail of a decrypt proposal",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd, nil
}
