package main

import (
	"github.com/ccpaging/cli/foo/commands"
	"os"
)

func main() {
	exitCode := commands.Run()
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}
