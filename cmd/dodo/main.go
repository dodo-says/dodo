package main

import "github.com/dodo-says/dodo/pkg/dodo/cmd"

func main() {
	rootCommand, err := cmd.NewRootCommand()
	if err != nil {
		panic(err)
	}
	_ = rootCommand.Execute()

}
