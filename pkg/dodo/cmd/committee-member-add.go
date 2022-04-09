package cmd

import (
	"context"
	"filippo.io/age/agessh"
	"github.com/dodo-says/dodo/pkg/committee"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

func NewCommitteeMemberAddCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	options := NewCommitteeMemberAddOptions("", "", "")
	cmd := &cobra.Command{
		Use:                   "add --committee-name <committee name> --public-key <path of public key> [-d <description> | --description <description>] <member-name> ",
		Short:                 "Add a new member to the committee",
		DisableFlagsInUseLine: true,
		Example: `
# Add a new member "alice" into the committee "dodo-says", with description "alice is trusted" and public key.
dodo committee member --committee-name dodo-says --public-key ./alice.pub --description "alice is trusted"  alice"
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			publicKeyContent, err := os.ReadFile(options.PublicKeyFilePath)
			if err != nil {
				return errors.Wrapf(err, "read public key file %s", options.PublicKeyFilePath)
			}
			_, err = agessh.ParseRecipient(string(publicKeyContent))
			if err != nil {
				return errors.Wrapf(err, "parse ssh public key file %s", options.PublicKeyFilePath)
			}
			memberName := args[0]
			member := committee.Member{
				Name:          memberName,
				CommitteeName: options.CommitteeName,
				Description:   options.Description,
				PublicKey:     publicKeyContent,
			}

			committeeService := BootstrapCommitteeService(globalOptions.StorageDir)
			err = committeeService.AddMemberToCommittee(ctx, options.CommitteeName, member)
			if err != nil {
				return errors.Wrapf(err, "add member %s to committee %s", memberName, options.CommitteeName)
			}

			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("missing required member name")
			}
			if len(args) > 1 {
				return errors.New("too many arguments, expected only one")
			}
			return nil
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
	}

	cmd.Flags().StringVar(&options.CommitteeName, "committee-name", "", "The name of the committee to add the member to")
	cmd.Flags().StringVarP(&options.Description, "description", "d", "", "The description of the member")
	cmd.Flags().StringVar(&options.PublicKeyFilePath, "public-key", "", "The public key of the member")

	err := cmd.MarkFlagRequired("committee-name")
	if err != nil {
		return nil, errors.Wrapf(err, "within command %s, mark flag %s required", cmd.Name(), "committee-name")
	}
	err = cmd.MarkFlagRequired("public-key")
	if err != nil {
		return nil, errors.Wrapf(err, "within command %s, mark flag %s required", cmd.Name(), "public-key")
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

type CommitteeMemberAddOptions struct {
	CommitteeName     string
	Description       string
	PublicKeyFilePath string
}

func NewCommitteeMemberAddOptions(committeeName string, description string, publicKeyFilePath string) *CommitteeMemberAddOptions {
	return &CommitteeMemberAddOptions{CommitteeName: committeeName, Description: description, PublicKeyFilePath: publicKeyFilePath}
}
