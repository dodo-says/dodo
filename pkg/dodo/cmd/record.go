package cmd

import "github.com/spf13/cobra"

func NewRecordCommand(globalOptions *GlobalOptions) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "record",
		Short: "Operations about record",
	}

	recordAddCommand, err := NewRecordAddCommand(globalOptions)
	if err != nil {
		return nil, err
	}
	cmd.AddCommand(recordAddCommand)

	recordListCommand, err := NewRecordListCommand(globalOptions)
	if err != nil {
		return nil, err
	}

	cmd.AddCommand(recordListCommand)
	cmd.AddCommand(NewRecordDecryptCommand(globalOptions))
	cmd.AddCommand(NewRecordClearCommand(globalOptions))

	recordGetEncryptedSliceCmd, err := NewRecordGetEncryptedSliceCmd(globalOptions)
	if err != nil {
		return nil, err
	}
	cmd.AddCommand(recordGetEncryptedSliceCmd)

	return cmd, nil
}
