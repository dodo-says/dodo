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
	}

	cmd.Flags().StringVar(&options.CommitteeName, "committee-name", "", "Committee name")
	err := cmd.MarkFlagRequired("committee-name")
	if err != nil {
		return nil, errors.Wrapf(err, "within command %s, mark flag %s as required", "list", "committee-name")
	}
	return cmd, nil
}

type CommitteeMemberListOptions struct {
	CommitteeName string
}

func NewCommitteeMemberListOptions(committeeName string) *CommitteeMemberListOptions {
	return &CommitteeMemberListOptions{CommitteeName: committeeName}
}
