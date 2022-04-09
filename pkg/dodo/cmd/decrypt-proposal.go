package cmd

import "github.com/spf13/cobra"

func NewDecryptProposalCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "decrypt-proposal",
		Short: "Committee proposal commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	proposalCreateCommand, err := NewDecryptProposalCreateCommand(globalOptions)
	if err != nil {
		return nil, err
	}
	cmd.AddCommand(proposalCreateCommand)
	cmd.AddCommand(NewDecryptProposalApproveCommand(globalOptions))
	proposalInspectCommand, err := NewDecryptProposalInspectCommand(globalOptions)
	if err != nil {
		return nil, err
	}
	cmd.AddCommand(proposalInspectCommand)

	proposalListCommand, err := NewDecryptProposalListCommand(globalOptions)
	if err != nil {
		return nil, err
	}
	cmd.AddCommand(proposalListCommand)

	return cmd, nil
}
