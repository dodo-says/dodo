package cmd

import (
	"context"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

func NewCommitteeMemberListCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	options := NewCommitteeMemberListOptions("")
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List committee members of a committee",
		Example: `
# List committee members in committee dodo
dodo committee-member list --committee-name dodo
`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			committeeService := BootstrapCommitteeService(globalOptions.StorageDir)
			members, err := committeeService.ListMemberOfCommittee(ctx, options.CommitteeName)
			if err != nil {
				return errors.Wrapf(err, "list committee members of %s", options.CommitteeName)
			}
			for _, member := range members {
				cmd.Println(member.Name)
			}
			return nil
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
	}

	cmd.Flags().StringVar(&options.CommitteeName, "committee-name", "", "Committee name")

	err := cmd.MarkFlagRequired("committee-name")
	if err != nil {
		return nil, errors.Wrapf(err, "within command %s, mark flag %s as required", cmd.Name(), "committee-name")
	}

	err = cmd.RegisterFlagCompletionFunc("committee-name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		ctx := context.TODO()
		committeeService := BootstrapCommitteeService(globalOptions.StorageDir)
		committees, err := committeeService.ListCommittee(ctx)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		var committeeNames []string
		for _, item := range committees {
			committeeNames = append(committeeNames, item.Name)
		}
		return committeeNames, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, errors.Wrapf(err, "within command %s, register flag completion for %s", cmd.Name(), "committee-name")
	}

	return cmd, nil
}

type CommitteeMemberListOptions struct {
	CommitteeName string
}

func NewCommitteeMemberListOptions(committeeName string) *CommitteeMemberListOptions {
	return &CommitteeMemberListOptions{CommitteeName: committeeName}
}
