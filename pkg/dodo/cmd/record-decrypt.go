package cmd

import "github.com/spf13/cobra"

func NewRecordDecryptCommand(globalOption *GlobalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypts a record",
		Long:  "Decrypts a record",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
