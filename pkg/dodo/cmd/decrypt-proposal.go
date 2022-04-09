package cmd

import "github.com/spf13/cobra"

func NewDecryptProposalCommand(globalOptions *GlobalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrypt-proposal",
		Short: "Committee proposal commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(NewDecryptProposalCreateCommand(globalOptions))
	cmd.AddCommand(NewDecryptProposalApproveCommand(globalOptions))

	return cmd
}
