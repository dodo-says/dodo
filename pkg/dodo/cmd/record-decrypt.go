package cmd

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strings"
)

func NewRecordDecryptCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	options := NewRecordDecryptOptions("")

	cmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypts a record",
		Args:  cobra.NoArgs,
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			recordUUID, err := uuid.Parse(options.RecordID)
			if err != nil {
				return errors.Wrapf(err, "invalid record id: %s", options.RecordID)
			}
			recordService := BootstrapRecordService(globalOptions.StorageDir)
			record, err := recordService.GetRecord(ctx, recordUUID)
			if err != nil {
				return errors.Wrapf(err, "get record with id %s", recordUUID)
			}
			proposalService := BootstrapProposalService(globalOptions.StorageDir)
			proposals, err := proposalService.ListDecryptProposalByRecordID(ctx, recordUUID)
			if err != nil {
				return errors.Wrapf(err, "list proposal by record id %s", recordUUID)
			}
			if len(proposals) == 0 {
				cmd.Println("No proposal found, you could use \"dodo decrypt-proposal create\" to create one")
				return nil
			}
			for _, proposal := range proposals {
				proposalApprovals, err := proposalService.ListDecryptProposalApprovalByProposalID(ctx, proposal.ProposalID)
				if err != nil {
					// best efforts, skip
					continue
				}

				if len(proposalApprovals) >= record.Threshold {
					decryptedRecord, err := proposalService.DecryptTheRecord(ctx, *record, proposal, proposalApprovals)
					if err != nil {
						// best efforts, skip
						continue
					}
					// decrypt succeed
					cmd.Println(fmt.Sprintf("Decrypted record: %s", decryptedRecord.PlaintextContent))
					return nil
				}
			}

			cmd.Println("No proposal has enough approvals, you should concat with other committee members for more approvals")
			cmd.Println("Available proposals:")
			for _, proposal := range proposals {
				approvals, err := proposalService.ListDecryptProposalApprovalByProposalID(ctx, proposal.ProposalID)
				if err != nil {
					// best efforts, skip
					continue
				}
				var approvedMembers []string
				for _, approval := range approvals {
					approvedMembers = append(approvedMembers, approval.Member)
				}
				cmd.Println(fmt.Sprintf("Proposal ID: %s, Reason: %s, threshould: %d, approved members: %s", proposal.ProposalID, proposal.Reason, record.Threshold, strings.Join(approvedMembers, ", ")))
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&options.RecordID, "record-id", "", "Record ID to decrypt")

	err := cmd.MarkFlagRequired("record-id")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s asrequired", "record-id")
	}

	err = cmd.RegisterFlagCompletionFunc("record-id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		recordService := BootstrapRecordService(globalOptions.StorageDir)
		records, err := recordService.ListRecords(context.TODO())
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var recordIDs []string
		for _, record := range records {
			recordIDs = append(recordIDs, record.ID.String())
		}
		return recordIDs, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, errors.Wrapf(err, "register flag completion for %s", "record-id")
	}

	return cmd, nil
}

type RecordDecryptOptions struct {
	RecordID string
}

func NewRecordDecryptOptions(recordID string) *RecordDecryptOptions {
	return &RecordDecryptOptions{RecordID: recordID}
}
