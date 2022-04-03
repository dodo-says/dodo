package cmd

import (
	"github.com/spf13/cobra"
)

func NewRecordClearCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Clear all records",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
