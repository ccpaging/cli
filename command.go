package cli

import (
	"os"
	"strings"
)

// CmdHandler is a handling function type for functions.
//
// Returned integer would be used as application exit status.
type CmdHandler func(args *Args) (exitCode int)

// Command represents a top-level application subcommand.
type Command struct {
	// Name is a [A-Za-z_0-9] identifier of up to 11 characters.
	//
	// Keep command names short, reasonable, catchy and
	// easy to type. At best, keep it a single word.
	//
	// Examples: build, list, install
	Name string

	// Brief is a short annotation of action command is capable of.
	//
	// Cli doesn't provide any limitations on the brief string
	// format, however it's highly recommended to keep it a single
	// lowercase phrase of 3-5 words without any punctuation marks.
	//
	// Example: compile packages and dependencies
	Brief string

	// Usage is a generic command use case, suggested by help.
	//
	// This string gets displayed on the usage line of command
	// help entry. You should NOT include command name itself.
	//
	// Example: [-o output] [-i] [build flags] [packages]
	Usage string

	// Help is a detailed command reference displayed after
	// the usage line and before the available flags block
	// of the help entry.
	//
	// Try to stick to the 80 character limit, so it looks fine
	// in the split terminal window.
	Help string

	// Division is the divsion displayed in help.
	Division string

	// Handling, I bet it's pretty straight-forward.
	Handle CmdHandler

	// Flags are command-line options.
	Flags []*Flag

	// Examples are annotated tips on command usage.
	Examples []*Example
}

// AddFlag does literally what its name says.
func (cmd *Command) AddFlag(newFlag *Flag) {
	cmd.Flags = append(cmd.Flags, newFlag)
}

// AddExample does exactly what its name says.
func (cmd *Command) AddExample(newExample *Example) {
	cmd.Examples = append(cmd.Examples, newExample)
}

// Run executes a command handler and returns corresponding exitcode.
func (cmd Command) Run(a *App) (exitCode int) {
	arguments := os.Args[1:]
	if len(arguments) > 0 && !strings.HasPrefix(arguments[0], "-") {
		// skip subcommand
		arguments = os.Args[2:]
	}
	ctx, err := newContext(a, cmd.Flags, arguments)
	if err != nil {
		a.printerr(err)
		os.Exit(1)
	}
	exitCode = cmd.Handle(ctx)
	return
}

// Topic is some sort of a concise wiki page.
type Topic struct {
	// Name is a [A-Za-z_0-9] identifier of up to 11 characters.
	//
	// Keep topic names short, reasonable, catchy and
	// easy to type. At best, keep it a single word.
	//
	// Examples: buildmode, packages, filetype
	Name string

	// Brief is a short annotation of the topic.
	//
	// Cli doesn't provide any limitations on the brief string
	// format, however it's highly recommended to keep it a single
	// lowercase phrase of 3-5 words without any punctuation marks.
	//
	// Example: description of package lists
	Brief string

	// Text is the actual topic content.
	//
	// Try to stick to the 80 character limit, so it looks fine
	// in the split terminal window.
	Text string
}

// Flag is an optional command-line option.
type Flag struct {
	// A flag label without the prefix (--, -, whatever).
	//
	// Flag names can't contain more than 11 alphanumeric characters.
	Name string

	// Usually the first letter of the name.
	//
	// Short names can't contain more than 3 alphanumeric characters.
	Short string

	// Default value (as string).
	DefValue string

	// Suggested use case, a generic example, showing
	// user how to use the flag.
	//
	// Example: --filter="token"
	Usage string

	// Help is displayed under the corresponding flag's
	// usage in the available commands section of help entry.
	//
	// Example: Limit tool output to tokens given.
	Help string
}

// Example is an annotated use case of the command.
type Example struct {
	// Usecase is a typical use of command.
	//
	// Make sure to omit application and command name here,
	// since Cli appends it by default.
	Usecase string

	// Be descriptive, but keep it under 3-5 sentences.
	Description string
}
