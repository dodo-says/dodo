package cmd

import (
	"context"
	"github.com/dodo-says/dodo/pkg/record"
	"github.com/dodo-says/dodo/pkg/share"
	"github.com/google/uuid"
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
			committeeService := BootstrapCommitteeService(globalOptions.StorageDir)

			members, err := committeeService.ListMemberOfCommittee(ctx, options.CommitteeName)
			if err != nil {
				return errors.Wrapf(err, "list members of committee %s", options.CommitteeName)
			}

			publicKeys := make(map[string][]byte)
			for _, member := range members {
				publicKeys[member.Name] = member.PublicKey
			}

			encryptedSlices, err := share.SplitThenEncrypt([]byte(options.Message), len(members), options.Threshold, publicKeys)
			if err != nil {
				return errors.Wrap(err, "encrypt message")
			}

			recordId := uuid.New()

			for _, slice := range encryptedSlices {
				err = recordService.AddEncryptedRecordSlice(ctx, record.EncryptedRecordSlice{
					ID:         uuid.New(),
					RecordID:   recordId,
					MemberName: slice.Name,
					Content:    slice.Content,
				})
				if err != nil {
					return errors.Wrap(err, "save encrypted record slice")
				}
			}

			err = recordService.AddRecord(ctx, record.Record{
				ID:            recordId,
				Description:   options.Description,
				CommitteeName: options.CommitteeName,
				Threshold:     options.Threshold,
			})
			if err != nil {
				return errors.Wrapf(err, "save record, description %s, committee %s, threshold %d", options.Description, options.CommitteeName, options.Threshold)
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
	err = cmd.MarkFlagDirname("threshold")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s as dirname", "threshold")
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
