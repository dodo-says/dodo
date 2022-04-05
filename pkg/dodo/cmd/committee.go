package cmd

import "github.com/spf13/cobra"

func NewCommitteeCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "committee",
		Short: "dodo committee commands",
		Args:  cobra.NoArgs,
	}

	// committee add
	cmd.AddCommand(NewCommitteeAddCommand(globalOptions))
	// committee get
	cmd.AddCommand(NewCommitteeGetCommand(globalOptions))
	// committee list
	cmd.AddCommand(NewCommitteeListCommand(globalOptions))
	// committee remove
	cmd.AddCommand(NewCommitteeRemoveCommand(globalOptions))

	return cmd, nil
}
