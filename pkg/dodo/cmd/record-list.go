package cmd

import "github.com/spf13/cobra"

func NewRecordListCommand(*GlobalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List records",
		Long:  "List records",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
