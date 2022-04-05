package cmd

import "github.com/spf13/cobra"

func NewRecordAddCommand(*GlobalOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new record",
		Long:  `Add a new record`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
