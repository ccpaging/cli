package main

import (
	"fmt"
	"github.com/ccpaging/cli"
	"strings"
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

		Handle: func(ctx *cli.Context) int {
			var separator string
			if sep, ok := ctx.ValueOf("separator"); ok {
				separator = sep
			}

			parts := strings.Split(separator, " ")
			if len(parts) > 1 {
				fmt.Println(strings.Join(parts[1:len(parts)], parts[0]))
			} else if len(parts) == 1 {
				fmt.Println(parts[0])
			}

			return 0
		},
	}

	demo.AddCommand(joinCmd)
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
