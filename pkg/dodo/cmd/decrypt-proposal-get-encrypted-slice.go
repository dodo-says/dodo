package cmd

import (
	"context"
	"filippo.io/age/armor"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewDecryptProposalGetEncryptedSliceCmd(globalOptions *GlobalOptions) (*cobra.Command, error) {
	options := NewRecordGetEncryptedSliceOptions("", "", true)

	cmd := &cobra.Command{
		Use:   "get-encrypted-slice",
		Short: "Get encrypted slice",
		Long:  `Get encrypted slice`,
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
				return errors.Wrapf(err, "get proposal %s", options.ProposalID)
			}
			recordUUID := proposal.RecordID
			slice, err := recordService.ListEncryptedRecordSliceByRecordIDAndMemberName(ctx, recordUUID, options.MemberName)
			if err != nil {
				return errors.Wrapf(err, "get encrypted slice, recordID: %s, memberName: %s", recordUUID, options.MemberName)
			}
			if options.Armored {
				armoredWriter := armor.NewWriter(cmd.OutOrStdout())
				_, err = armoredWriter.Write(slice.Content)
				if err != nil {
					return errors.Wrap(err, "write armored data")
				}
				err = armoredWriter.Close()
				if err != nil {
					return errors.Wrap(err, "close age armored writer")
				}
			} else {
				_, err = cmd.OutOrStdout().Write(slice.Content)
				if err != nil {
					return errors.Wrap(err, "write raw data")
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&options.ProposalID, "proposal-id", "r", "", "proposal id")
	cmd.Flags().StringVarP(&options.MemberName, "member-name", "m", "", "member name")
	cmd.Flags().BoolVar(&options.Armored, "armored", true, "armored")

	err := cmd.MarkFlagRequired("proposal-id")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s required", "proposal-id")
	}
	err = cmd.MarkFlagRequired("member-name")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s required", "member-name")
	}

	err = cmd.RegisterFlagCompletionFunc("proposal-id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		ctx := context.TODO()
		proposalService := BootstrapProposalService(globalOptions.StorageDir)
		proposals, err := proposalService.ListDecryptProposal(ctx)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var ids []string
		for _, proposal := range proposals {
			ids = append(ids, proposal.ProposalID.String())
		}
		return ids, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, errors.Wrap(err, "register flag completion")
	}
	err = cmd.RegisterFlagCompletionFunc("member-name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		ctx := context.TODO()
		proposalUUID, err := uuid.Parse(options.ProposalID)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		proposalService := BootstrapProposalService(globalOptions.StorageDir)
		proposal, err := proposalService.GetDecryptProposal(ctx, proposalUUID)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		recordService := BootstrapRecordService(globalOptions.StorageDir)
		slices, err := recordService.ListEncryptedRecordSlicesByRecordID(ctx, proposal.RecordID)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var names []string
		for _, slice := range slices {
			names = append(names, slice.MemberName)
		}
		return names, cobra.ShellCompDirectiveNoFileComp
	})

	return cmd, nil
}

type RecordGetEncryptedSliceOptions struct {
	ProposalID string
	MemberName string
	Armored    bool
}

func NewRecordGetEncryptedSliceOptions(proposalID string, memberName string, armored bool) *RecordGetEncryptedSliceOptions {
	return &RecordGetEncryptedSliceOptions{ProposalID: proposalID, MemberName: memberName, Armored: armored}
}
