package cmd

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCommitteeGetCommand(globalOptions *GlobalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "get [flags] committee-name",
		DisableFlagsInUseLine: true,
		Short:                 "Get committee by name",
		Example: `
# get committee with name "dodo"
dodo committee get dodo
`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("missing required committee name")
			}
			if len(args) > 1 {
				return errors.New("too many arguments, expected only one")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			committeeName := args[0]
			committeeService := BootstrapCommitteeService(globalOptions.StorageDir)
			committee, err := committeeService.GetCommittee(ctx, committeeName)
			if err != nil {
				return errors.Wrapf(err, "get committee %s", committeeName)
			}
			cmd.Printf("Name: %s\nDescription: %s\n", committee.Name, committee.Description)
			return nil
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) > 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			ctx := context.TODO()
			committeeService := BootstrapCommitteeService(globalOptions.StorageDir)
			committees, err := committeeService.ListCommittee(ctx)
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}

			var names []string
			for _, committee := range committees {
				names = append(names, committee.Name)
			}
			return names, cobra.ShellCompDirectiveNoFileComp
		},
	}

	return cmd
}
