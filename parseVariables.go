package cli

import (
	"fmt"
	"strings"
)

func parseVariables(beStrict bool, flags []*Flag, argv []string) (map[string]string, error) {
	vars := make(map[string]string, 0)
	for i := 0; i < len(argv); i++ {
		argument := argv[i]

		if !strings.HasPrefix(argument, "-") {
			if beStrict {
				return nil, fmt.Errorf(`No option name before argument %s`, argument)
			}
			continue
		}

		argument = strings.TrimLeft(argument, "-")

		parts := strings.SplitN(argument, "=", 2)

		name := parts[0]

		var flag *Flag
		for _, f := range flags {
			if f.Name == name || f.Short == name {
				flag = f
				break
			}
		}
		if flag == nil {
			if beStrict {
				return nil, fmt.Errorf(`option -%s does not exist`, name)
			}
			flag = &Flag{Name: name}
		}

		value := ""
		if len(parts) > 1 {
			value = parts[1]
		}

		for i+1 < len(argv) {
			if strings.HasPrefix(argv[i+1], "-") {
				break
			}
			value += " " + argv[i+1]
			i++
		}

		if len(parts) == 1 && len(value) == 0 {
			vars[flag.Name] = ""
			continue
		}

		vars[flag.Name] = strings.TrimLeft(value, " ")
	}

	return vars, nil
}
