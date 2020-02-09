package commands

import (
	"github.com/ccpaging/cli"
)

var App = &cli.App{
	Name:    "demo",
	Brief:   "Demo is a funky demonstation of Cli capabilities.",
	Version: "stable",
}

func Run() (exitCode int) {
	return App.Run()
}
