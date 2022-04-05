package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewRootCommand() *cobra.Command {
	globalOptions := &GlobalOptions{}
	rootCommand := &cobra.Command{
		Use:          "dodo",
		Short:        "dodo is the tool as the PoC of shamir secret sharing with co-share committee",
		SilenceUsage: true,
	}

	rootCommand.PersistentFlags().StringVar(&globalOptions.StorageDir, "storage-path", mustDefaultStorageDir(), "The path to the storage")
	rootCommand.AddCommand(NewCommitteeCommand(globalOptions))
	rootCommand.AddCommand(NewRecordCommand(globalOptions))

	return rootCommand
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
