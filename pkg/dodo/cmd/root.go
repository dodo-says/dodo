package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewRootCommand() (*cobra.Command, error) {
	globalOptions := &GlobalOptions{}
	rootCommand := &cobra.Command{
		Use:          "dodo",
		Short:        "dodo is the tool as the PoC of shamir secret sharing with co-share committee",
		SilenceUsage: true,
	}

	rootCommand.PersistentFlags().StringVar(&globalOptions.StorageDir, "storage-path", mustDefaultStorageDir(), "The path to the storage")

	// dodo committee
	committeeCommand, err := NewCommitteeCommand(globalOptions)
	if err != nil {
		return nil, err
	}
	rootCommand.AddCommand(committeeCommand)

	// dodo committee-member
	committeeMemberCommand, err := NewCommitteeMemberCommand(globalOptions)
	if err != nil {
		return nil, err
	}
	rootCommand.AddCommand(committeeMemberCommand)

	// dodo committee-proposal
	rootCommand.AddCommand(NewCommitteeProposalCommand(globalOptions))

	// dodo record
	recordCommand, err := NewRecordCommand(globalOptions)
	if err != nil {
		return nil, err
	}
	rootCommand.AddCommand(recordCommand)

	return rootCommand, nil
}

type GlobalOptions struct {
	StorageDir string
}

func mustDefaultStorageDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/.dodo", home)
}
