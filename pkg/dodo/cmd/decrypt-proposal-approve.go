package cmd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/dodo-says/dodo/pkg/proposal"
	"github.com/dodo-says/dodo/pkg/share"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io/ioutil"
)

func NewDecryptProposalApproveCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	options := NewDecryptProposalApproveOptions("")

	var cmd = &cobra.Command{
		Use:   "approve",
		Short: "Approve a decrypt proposal",
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			proposalUUID, err := uuid.Parse(options.ProposalID)
			if err != nil {
				return errors.Wrapf(err, "invalid proposal id: %s", options.ProposalID)

			}
			proposalService := BootstrapProposalService(globalOptions.StorageDir)
			// make sure that the proposal exists
			_, err = proposalService.GetDecryptProposal(context.Background(), proposalUUID)
			if err != nil {
				return errors.Wrapf(err, "get proposal: %s", options.ProposalID)
			}

			jsonInBytes, err := ioutil.ReadAll(cmd.InOrStdin())
			payload := share.Payload{
				MemberName:  "",
				SliceBase64: "",
			}
			err = json.Unmarshal(jsonInBytes, &payload)
			if err != nil {
				return errors.Wrap(err, "unmarshal payload")
			}
			plaintextShamirSlice, err := base64.StdEncoding.DecodeString(payload.SliceBase64)
			if err != nil {
				return errors.Wrap(err, "decode slice")
			}
			if err != nil {
				return errors.Wrap(err, "read plaintext slice from stdin")
			}
			err = proposalService.CreateDecryptProposalApproval(ctx, proposal.DecryptProposalApproval{
				ProposalID:     proposalUUID,
				Member:         payload.MemberName,
				PlaintextSlice: plaintextShamirSlice,
			})
			if err != nil {
				return errors.Wrapf(err, "create proposal approval for proposal: %s", options.ProposalID)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&options.ProposalID, "proposal-id", "", "Proposal ID")

	err := cmd.MarkFlagRequired("proposal-id")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s as required", "proposal-id")
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

type DecryptProposalApproveOptions struct {
	ProposalID string
}

func NewDecryptProposalApproveOptions(proposalID string) *DecryptProposalApproveOptions {
	return &DecryptProposalApproveOptions{ProposalID: proposalID}
}
