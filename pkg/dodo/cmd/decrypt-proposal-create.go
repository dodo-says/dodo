package cmd

import (
	"context"
	"github.com/dodo-says/dodo/pkg/proposal"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewDecryptProposalCreateCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	options := NewDecryptProposalCreateOptions("", "")
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new decrypt proposal",
		Example: `
# Create a new decrypt proposal for record 27808677-614e-4339-a8fb-7d4778fdf9ce with the reason "I think this message contains something really important"
dodo decrypt-proposal create --record-id 27808677-614e-4339-a8fb-7d4778fdf9ce --reason "I think this message contains something really important"
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			proposalService := BootstrapProposalService(globalOptions.StorageDir)
			recordService := BootstrapRecordService(globalOptions.StorageDir)

			recordUUID, err := uuid.Parse(options.RecordID)
			if err != nil {
				return errors.Wrapf(err, "invalid record id %s", options.RecordID)
			}
			// record must exist
			_, err = recordService.GetRecord(ctx, recordUUID)
			if err != nil {
				return errors.Wrapf(err, "get record %s", options.RecordID)
			}
			err = proposalService.CreateDecryptProposal(ctx, proposal.DecryptProposal{
				ProposalID: uuid.New(),
				RecordID:   recordUUID,
				Reason:     options.Reason,
			})
			if err != nil {
				return errors.Wrap(err, "create decrypt proposal")
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&options.RecordID, "record-id", "", "The ID of the record to decrypt")
	cmd.Flags().StringVar(&options.Reason, "reason", "", "The reason for why you want to decrypt this record")

	err := cmd.MarkFlagRequired("record-id")
	if err != nil {
		return nil, err
	}
	err = cmd.MarkFlagRequired("reason")
	if err != nil {
		return nil, err
	}

	err = cmd.RegisterFlagCompletionFunc("record-id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		ctx := context.TODO()
		recordService := BootstrapRecordService(globalOptions.StorageDir)
		records, err := recordService.ListRecords(ctx)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var recordIds []string
		for _, record := range records {
			recordIds = append(recordIds, record.ID.String())
		}
		return recordIds, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, errors.Wrapf(err, "registering flag completion for %s", "record-id")
	}
	err = cmd.RegisterFlagCompletionFunc("reason", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, errors.Wrapf(err, "registering flag completion for %s", "reason")
	}

	return cmd, nil
}

type DecryptProposalCreateOptions struct {
	RecordID string
	Reason   string
}

func NewDecryptProposalCreateOptions(recordID string, reason string) *DecryptProposalCreateOptions {
	return &DecryptProposalCreateOptions{RecordID: recordID, Reason: reason}
}
