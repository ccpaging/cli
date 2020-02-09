package commands

import (
	"fmt"
	"github.com/ccpaging/cli"
	"strings"
)

var JoinCmd = &cli.Command{
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
		sep, ok := ctx.ValueOf("separator")
		if !ok {
			fmt.Println("separator not specified")

			return 1
		}

		parts := strings.Split(sep, " ")
		if len(parts) > 0 {
			if len(parts) > 1 {
				fmt.Println(strings.Join(parts[1:len(parts)], parts[0]))
			} else {
				fmt.Println(parts[0])
			}
		}

		return 0
	},
}

func init() {
	App.AddCommand(JoinCmd)
}
