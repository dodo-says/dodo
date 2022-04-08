package cmd

import (
	"context"
	"github.com/dodo-says/dodo/pkg/committee"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCommitteeAddCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	options := NewCommitteeAddOptions("")

	cmd := &cobra.Command{
		Use:                   "add [-d <description> | --description <description>] committee-name",
		DisableFlagsInUseLine: true,
		Short:                 "Add a new committee",
		Example: `
# create a new committee with the name "Knights of the Round Table" and description "KOTRT want to share something secret"
dodo committee add -d "KOTRT want to share something secret" "Knights of the Round Table"

# create a new committee with the name "dodo-says" and without any description.
dodo committee add "dodo-says"
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
			committeeService := BootstrapCommitteeService(globalOptions.StorageDir)
			newCommittee := committee.NewCommittee(args[0], options.Description)
			err := committeeService.AddCommittee(ctx, *newCommittee)
			if err != nil {
				return errors.Wrap(err, "create committee")
			}
			return nil
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
	}

	cmd.Flags().StringVarP(&options.Description, "description", "d", "", "Optional description for this committee")
	err := cmd.RegisterFlagCompletionFunc("description", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, errors.Wrapf(err, "register completion for flag %s", "description")
	}

	return cmd, nil
}

// CommitteeAddOptions is the options for the committee add command.
type CommitteeAddOptions struct {
	// Description is the optional description for this committee.
	Description string
}

func NewCommitteeAddOptions(description string) *CommitteeAddOptions {
	return &CommitteeAddOptions{Description: description}
}
