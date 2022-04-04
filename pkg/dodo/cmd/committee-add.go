package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCommitteeAddCommand() *cobra.Command {
	options := NewCommitteeAddOptions("", "")

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

			return nil
		},
	}

	cmd.Flags().StringVarP(&options.Description, "description", "d", "", "Optional description for this committee")

	return cmd
}

// CommitteeAddOptions is the options for the committee add command.
type CommitteeAddOptions struct {
	// CommitteeName is the required name for this committee.
	CommitteeName string
	// Description is the optional description for this committee.
	Description string
}

func NewCommitteeAddOptions(committeeName string, description string) *CommitteeAddOptions {
	return &CommitteeAddOptions{CommitteeName: committeeName, Description: description}
}
