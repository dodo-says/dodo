package cmd

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strings"
)

func NewDecryptProposalInspectCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	options := NewDecryptProposalInspectOptions("")
	cmd := &cobra.Command{
		Use:   "inspect",
		Short: "inspect the detail of a decrypt proposal",
		Args:  cobra.NoArgs,
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			proposalUUID, err := uuid.Parse(options.ProposalID)
			if err != nil {
				return errors.Wrapf(err, "invalid proposal id %s", options.ProposalID)
			}

			proposalService := BootstrapProposalService(globalOptions.StorageDir)
			recordService := BootstrapRecordService(globalOptions.StorageDir)

			proposal, err := proposalService.GetDecryptProposal(ctx, proposalUUID)
			if err != nil {
				return errors.Wrapf(err, "get proposal by id  %s", options.ProposalID)
			}
			record, err := recordService.GetRecord(ctx, proposal.RecordID)
			if err != nil {
				return errors.Wrapf(err, "get record by id  %s", proposal.RecordID)
			}
			approvals, err := proposalService.ListDecryptProposalApprovalByProposalID(ctx, proposalUUID)
			if err != nil {
				return errors.Wrapf(err, "list approval by proposal id %s", options.ProposalID)
			}
			var approvedMembers []string
			for _, approval := range approvals {
				approvedMembers = append(approvedMembers, approval.Member)
			}

			cmd.Printf("Proposal ID: %s\n", proposal.ProposalID.String())
			cmd.Printf("Proposal Reason: %s\n", proposal.Reason)
			cmd.Printf("Record ID: %s\n", proposal.RecordID.String())
			cmd.Printf("Record Description: %s\n", record.Description)
			cmd.Printf("Committee: %s\n", record.CommitteeName)
			cmd.Printf("Approve Committee Members: %s\n", strings.Join(approvedMembers, ", "))

			return nil
		},
	}

	cmd.Flags().StringVar(&options.ProposalID, "proposal-id", "", "the id of proposal to inspect")
	err := cmd.MarkFlagRequired("proposal-id")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s as required failed", "proposal-id")
	}

	err = cmd.RegisterFlagCompletionFunc("proposal-id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		ctx := context.TODO()
		proposalService := BootstrapProposalService(globalOptions.StorageDir)
		proposals, err := proposalService.ListDecryptProposal(ctx)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var proposalIDs []string
		for _, proposal := range proposals {
			proposalIDs = append(proposalIDs, proposal.ProposalID.String())
		}
		return proposalIDs, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, errors.Wrapf(err, "register flag completion for %s", "proposal-id")
	}

	return cmd, nil
}

type DecryptProposalInspectOptions struct {
	ProposalID string
}

func NewDecryptProposalInspectOptions(proposalID string) *DecryptProposalInspectOptions {
	return &DecryptProposalInspectOptions{ProposalID: proposalID}
}
