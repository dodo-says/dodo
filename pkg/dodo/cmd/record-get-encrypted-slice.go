package cmd

import (
	"context"
	"filippo.io/age/armor"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewRecordGetEncryptedSliceCmd(globalOptions *GlobalOptions) (*cobra.Command, error) {
	options := NewRecordGetEncryptedSliceOptions("", "", true)

	cmd := &cobra.Command{
		Use:   "get-encrypted-slice",
		Short: "Get encrypted slice",
		Long:  `Get encrypted slice`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			recordService := BootstrapRecordService(globalOptions.StorageDir)
			recordUUID, err := uuid.Parse(options.RecordID)
			if err != nil {
				return errors.Wrapf(err, "prase record id %s", options.RecordID)
			}
			slice, err := recordService.ListEncryptedRecordSliceByRecordIDAndMemberName(ctx, recordUUID, options.MemberName)
			if err != nil {
				return errors.Wrapf(err, "get encrypted slice, recordID: %s, memberName: %s", options.RecordID, options.MemberName)
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

	cmd.Flags().StringVarP(&options.RecordID, "record-id", "r", "", "record id")
	cmd.Flags().StringVarP(&options.MemberName, "member-name", "m", "", "member name")
	cmd.Flags().BoolVar(&options.Armored, "armored", true, "armored")

	err := cmd.MarkFlagRequired("record-id")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s required", "record-id")
	}
	err = cmd.MarkFlagRequired("member-name")
	if err != nil {
		return nil, errors.Wrapf(err, "mark flag %s required", "member-name")
	}

	err = cmd.RegisterFlagCompletionFunc("record-id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		ctx := context.TODO()
		recordService := BootstrapRecordService(globalOptions.StorageDir)
		records, err := recordService.ListRecords(ctx)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var ids []string
		for _, record := range records {
			ids = append(ids, record.ID.String())
		}
		return ids, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, errors.Wrap(err, "register flag completion")
	}
	err = cmd.RegisterFlagCompletionFunc("member-name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		ctx := context.TODO()
		recordService := BootstrapRecordService(globalOptions.StorageDir)
		recordUUID, err := uuid.Parse(options.RecordID)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		slices, err := recordService.ListEncryptedRecordSlicesByRecordID(ctx, recordUUID)
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
	RecordID   string
	MemberName string
	Armored    bool
}

func NewRecordGetEncryptedSliceOptions(recordID string, memberName string, armored bool) *RecordGetEncryptedSliceOptions {
	return &RecordGetEncryptedSliceOptions{RecordID: recordID, MemberName: memberName, Armored: armored}
}
