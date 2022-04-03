package cmd

import "github.com/spf13/cobra"

func NewRecordCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "record",
		Short: "Operations about record",
	}

	cmd.AddCommand(NewRecordAddCommand())
	cmd.AddCommand(NewRecordListCommand())
	cmd.AddCommand(NewRecordDecryptCommand())
	cmd.AddCommand(NewRecordClearCommand())

	return cmd
}
