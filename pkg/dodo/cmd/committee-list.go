package cmd

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCommitteeListCommand(globalOptions *GlobalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all committees",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			committeeService := BootstrapCommitteeService(globalOptions.StorageDir)
			committees, err := committeeService.ListCommittee(ctx)
			if err != nil {
				return errors.Wrap(err, "list committees")
			}
			for _, committee := range committees {
				cmd.Println(committee.Name)
			}
			return nil
		},
	}
	return cmd
}
