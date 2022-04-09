package cmd

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewDecryptProposalListCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list decrypt proposals",
		Args:  cobra.NoArgs,
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			proposalService := BootstrapProposalService(globalOptions.StorageDir)
			proposals, err := proposalService.ListDecryptProposal(ctx)
			if err != nil {
				return errors.Wrapf(err, "list decrypt proposals")
			}
			cmd.Println("ProposalID\tRecordID\tReason")
			for _, proposal := range proposals {
				cmd.Printf("%s\t%s\t%s\n", proposal.ProposalID, proposal.RecordID, proposal.Reason)
			}
			return nil
		},
	}
	return cmd, nil
}
