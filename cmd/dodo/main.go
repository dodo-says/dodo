package main

import "github.com/dodo-says/dodo/pkg/dodo/cmd"

func main() {
	_ = cmd.NewRootCommand().Execute()
}
