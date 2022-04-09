package cmd

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewRecordListCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	options := NewRecordListOptions("")
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List records",
		Long:  "List records",
		Args:  cobra.NoArgs,
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			recordService := BootstrapRecordService(globalOptions.StorageDir)
			records, err := recordService.ListRecordsByCommittee(ctx, options.CommitteeName)
			if err != nil {
				return errors.Wrapf(err, "list records by committee %s", options.CommitteeName)
			}
			cmd.Println(fmt.Sprintf("%s\t%s\t%s\t%s", "ID", "Description", "Committee", "Threshold"))
			for _, record := range records {
				cmd.Println(fmt.Sprintf("%s\t%s\t%s\t%d", record.ID, record.Description, record.CommitteeName, record.Threshold))
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&options.CommitteeName, "committee-name", "", "Committee name")

	err := cmd.MarkFlagRequired("committee-name")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s as required", "committee-name")
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
		return nil, errors.Wrapf(err, "register flag completion for %s", "committee-name")
	}
	return cmd, nil
}

type RecordListOptions struct {
	CommitteeName string
}

func NewRecordListOptions(committeeName string) *RecordListOptions {
	return &RecordListOptions{CommitteeName: committeeName}
}
