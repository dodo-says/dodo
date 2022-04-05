package cmd

import "github.com/spf13/cobra"

func NewCommitteeMemberCommand(globalOption *GlobalOptions) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "committee-member",
		Short: "dodo committee member commands",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	memberAddCommand, err := NewCommitteeMemberAddCommand(globalOption)
	if err != nil {
		return nil, err
	}
	cmd.SilenceUsage = true

	cmd.AddCommand(memberAddCommand)
	cmd.AddCommand(NewCommitteeMemberRemoveCommand(globalOption))
	cmd.AddCommand(NewCommitteeMemberListCommand(globalOption))

	return cmd, nil
}
