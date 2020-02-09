package cli

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	Stdout io.Writer = os.Stdout
	Stderr io.Writer = os.Stderr
)

// Cli is a main CLI instance.
//
// By default, Cli provides its own implementation of version
// command, but it will use "version" command instead if you
// provide one.
type App struct {
	Name    string // `go`
	Brief   string // `Go is a tool for managing Go source code.`
	Version string // `1.5`
	Strict  bool   // default is false

	Root     *Command
	Commands []*Command
	Topics   []*Topic
}

// New constructs a new CLI application with a given name.
// In case of an empty name it will panic.
func NewApp(name string) *App {
	if name == "" {
		panic("can't construct an app without a name")
	}

	return &App{Name: name, Strict: true}
}

func (a *App) println(stuff ...interface{}) {
	fmt.Fprintln(Stdout, stuff...)
}

func (a *App) printf(format string, stuff ...interface{}) {
	fmt.Fprintf(Stdout, format, stuff...)
}

func (a *App) printerr(err ...interface{}) {
	for _, each := range err {
		fmt.Fprintln(Stderr, a.Name+":", each)
	}
}

func (a *App) commandByName(name string) *Command {
	for i, command := range a.Commands {
		if command.Name == name {
			return a.Commands[i]
		}
	}

	return nil
}

func (a *App) topicByName(name string) *Topic {
	for i, topic := range a.Topics {
		if topic.Name == name {
			return a.Topics[i]
		}
	}

	return nil
}

// AddCommand does literally what its name says.
func (a *App) AddCommand(command *Command) {
	a.Commands = append(a.Commands, command)
}

// AddTopic does literally what its name says.
func (a *App) AddTopic(topic *Topic) {
	a.Topics = append(a.Topics, topic)
}

// SuggestionsFor provides suggestions for the typedName.
func (a *App) SuggestionsFor(typedName string) []string {
	minimumDistance := 2
	suggestions := []string{}
	for _, cmd := range a.Commands {
		ld := levenshteinDistance(typedName, cmd.Name, true)
		hasPrefix := strings.HasPrefix(strings.ToLower(cmd.Name), strings.ToLower(typedName))
		if ld <= minimumDistance || hasPrefix {
			suggestions = append(suggestions, cmd.Name)
		}
	}
	return suggestions
}

// Run executes a a.
//
// Take a note, Run panics if len(os.Args) < 1
func (a *App) Run() int {
	if len(os.Args) < 1 {
		panic("shell-provided arguments are not present")
	}
	arguments := os.Args[1:]

	// $ program
	// $ program -flag
	//           ^ no subcommand
	if len(arguments) == 0 || ((len(arguments) > 0) && strings.HasPrefix(arguments[0], "-")) {
		if a.Root != nil {
			return a.Root.Run(a)
		}

		a.println(a.globalHelp())
		return 0
	}

	subcommandName := arguments[0]
	subcommand := a.commandByName(subcommandName)

	if subcommandName == "help" {
		// $ program help
		//           ^ one argument
		if len(arguments) <= 1 {
			if a.Root != nil {
				a.println(a.commandHelp(a.Root))
			} else {
				a.println(a.globalHelp())
			}
			return 0
		}

		command := a.commandByName(arguments[1])
		if command != nil {
			a.println(a.commandHelp(command))
			return 0
		}

		topic := a.topicByName(arguments[1])
		if topic != nil {
			a.println(topic.Text)
			return 0
		}

		a.printerr("no such command or help topic")
		os.Exit(1)
	}

	if subcommandName == "version" {
		if subcommand != nil {
			return subcommand.Run(a)
		}

		a.printf("%s version %s\n", a.Name, a.Version)
		return 0
	}

	if subcommand != nil {
		return subcommand.Run(a)
	}

	a.printerr("unknown subcommand \"" + subcommandName + "\"\n")
	if suggestions := a.SuggestionsFor(subcommandName); len(suggestions) > 0 {
		a.println("Did you mean this?")
		for _, s := range suggestions {
			a.printf("\t%v\n", s)
		}
		a.println("")
	}
	os.Exit(1)

	return 1
}
