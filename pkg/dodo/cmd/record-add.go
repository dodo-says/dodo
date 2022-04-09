package cmd

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewRecordAddCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	options := NewRecordAddOptions("", 0, "", "")

	cmd := &cobra.Command{
		Use:                   "add --committee-name <committee-name> {-t <threshold> | --threshold <threshold>} --message <message> [-d <description> | --description <description>]",
		DisableFlagsInUseLine: true,
		Short:                 "Add a new record",
		Example: `
# Add a record "dodo helps dodo" under committee "dodo", at least 2 committee approve are required to decrypt this record
dodo record add --committee-name dodo --threshold 2 --message "dodo helps dodo"

# Add a record "STRRL is a dodo" under committee "dodo", with a plaintext description "STRRL's introduction", at least 3 committee approve are required to decrypt this record
dodo record add --committee-name dodo -t 3 --message "STRRL is a dodo" -d "STRRL's introduce"
`,
		Args: cobra.NoArgs,
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			recordService := BootstrapRecordService(globalOptions.StorageDir)

			record, slices, err := recordService.BuildRecord(ctx, options.Message, options.Description, options.CommitteeName, options.Threshold)
			if err != nil {
				return errors.Wrapf(err, "build record for committee %s, description %s, threshold %d", options.CommitteeName, options.Description, options.Threshold)
			}
			err = recordService.AddRecord(ctx, *record)
			if err != nil {
				return errors.Wrap(err, "save record")
			}

			for _, slice := range slices {
				err = recordService.AddEncryptedRecordSlice(ctx, slice)
				if err != nil {
					return errors.Wrap(err, "save encrypted slice")
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&options.Description, "description", "d", "", "Plaintext description of the record, it would not be encrypted")
	cmd.Flags().StringVar(&options.CommitteeName, "committee-name", "", "Name of the committee")
	cmd.Flags().StringVar(&options.Message, "message", "", "Message of the record, it would be encrypted")
	cmd.Flags().IntVarP(&options.Threshold, "threshold", "t", 0, "Threshold means: at least this number of committee members must approve to decrypt the record")

	err := cmd.MarkFlagRequired("committee-name")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s as required", "committee-name")
	}
	err = cmd.MarkFlagRequired("message")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s as required", "message")
	}
	err = cmd.MarkFlagRequired("threshold")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s as dirname", "threshold")
	}

	err = cmd.RegisterFlagCompletionFunc("description", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, errors.Wrapf(err, "register flag completion function for %s", "description")
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
		return nil, errors.Wrapf(err, "register flag completion function for %s", "committee-name")
	}
	err = cmd.RegisterFlagCompletionFunc("message", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, errors.Wrapf(err, "register flag completion function for %s", "message")
	}
	err = cmd.RegisterFlagCompletionFunc("threshold", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, errors.Wrapf(err, "register flag completion function for %s", "threshold")
	}
	return cmd, nil
}

type RecordAddOptions struct {
	CommitteeName string
	Threshold     int
	Message       string
	Description   string
}

func NewRecordAddOptions(committeeName string, threshold int, message string, description string) *RecordAddOptions {
	return &RecordAddOptions{CommitteeName: committeeName, Threshold: threshold, Message: message, Description: description}
}
