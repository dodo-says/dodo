package cmd

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "dodo",
		Short: "dodo is the tool as the PoC of shamir secret sharing with co-share committers",
	}
	rootCommand.AddCommand(NewCommitteeCommand())
	rootCommand.AddCommand(NewRecordCommand())
	return rootCommand
}
