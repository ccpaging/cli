package main

import (
	"fmt"
	"github.com/ccpaging/cli"
)

func main() {
	demo := cli.NewApp("demo")
	demo.Brief = "Demo is a funky demonstation of cli capabilities."
	demo.Version = "stable"

	joinCmd := &cli.Command{
		Name:  "join",
		Brief: "merges the strings given",
		Usage: `[-s=] "a few" distinct strings`,
		Help:  `Lorem ipsum dolor sit amet amet sit todor...`,

		Flags: []*cli.Flag{
			&cli.Flag{
				Name:  "separator",
				Short: "s",
				Usage: `--separator="."`,
				Help:  `Put some separating string between all the strings given.`,
			},
		},

		Examples: []*cli.Example{
			&cli.Example{
				Usecase:     `-s . "google" "com"`,
				Description: `Results in "google.com"`,
			},
		},

		Handle: Example_handler,
	}

	joinCmd1 := &cli.Command{
		Name:  "join1",
		Brief: "merges the strings given",
		Usage: `[-s=] "a few" distinct strings`,
		Help:  `Lorem ipsum dolor sit amet amet sit todor...`,

		Division: "join",

		Flags: []*cli.Flag{
			&cli.Flag{
				Name:  "separator",
				Short: "s",
				Usage: `--separator="."`,
				Help:  `Put some separating string between all the strings given.`,
			},
		},

		Examples: []*cli.Example{
			&cli.Example{
				Usecase:     `-s . "google" "com"`,
				Description: `Results in "google.com"`,
			},
		},

		Handle: Example_handler,
	}

	joinCmd2 := &cli.Command{
		Name:  "join2",
		Brief: "merges the strings given",
		Usage: `[-s=] "a few" distinct strings`,
		Help:  `Lorem ipsum dolor sit amet amet sit todor...`,

		Flags: []*cli.Flag{
			&cli.Flag{
				Name:  "separator",
				Short: "s",
				Usage: `--separator="."`,
				Help:  `Put some separating string between all the strings given.`,
			},
		},

		Examples: []*cli.Example{
			&cli.Example{
				Usecase:     `-s . "google" "com"`,
				Description: `Results in "google.com"`,
			},
		},

		Handle: Example_handler,
	}

	demo.AddCommand(joinCmd)
	// Division displayed in help
	demo.AddCommand(joinCmd1)
	demo.AddCommand(joinCmd2)
	demo.Run()
}

// Handler accepts a cli.Context object and returns an exitcode integer.
func Example_handler(ctx *cli.Context) int {
	name, ok := ctx.ValueOf("name")
	if !ok {
		fmt.Println("name not specified")

		return 1
	}

	// argument `name` parsed
	fmt.Println(name)

	return 0
}
