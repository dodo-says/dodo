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

	recordDecryptCommand, err := NewRecordDecryptCommand(globalOptions)
	if err != nil {
		return nil, err
	}
	cmd.AddCommand(recordDecryptCommand)

	cmd.AddCommand(NewRecordClearCommand(globalOptions))

	return cmd, nil
}
