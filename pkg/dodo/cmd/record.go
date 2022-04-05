package cmd

import "github.com/spf13/cobra"

func NewRecordCommand(globalOptions *GlobalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "record",
		Short: "Operations about record",
	}

	cmd.AddCommand(NewRecordAddCommand(globalOptions))
	cmd.AddCommand(NewRecordListCommand(globalOptions))
	cmd.AddCommand(NewRecordDecryptCommand(globalOptions))
	cmd.AddCommand(NewRecordClearCommand(globalOptions))

	return cmd
}
