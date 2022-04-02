package main

import "github.com/dodo-says/dodo/pkg/dodo/cmd"

func main() {
	err := cmd.NewRootCommand().Execute()
	if err != nil {
		panic(err)
	}
}
