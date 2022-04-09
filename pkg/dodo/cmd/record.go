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
	cmd.AddCommand(NewRecordListCommand(globalOptions))
	cmd.AddCommand(NewRecordDecryptCommand(globalOptions))
	cmd.AddCommand(NewRecordClearCommand(globalOptions))

	return cmd, nil
}
