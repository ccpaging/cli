package cli

import (
	"strconv"
)

// Args is a set of arguments and options of command call.
type Args struct {
	app  *App
	vars map[string]string
}

func newContext(a *App, flags []*Flag, argv []string) (*Args, error) {
	vars, err := parseVariables(a.Strict, flags, argv)
	if err != nil {
		return nil, err
	}

	c := &Args{
		app:  a,
		vars: vars,
	}
	return c, nil
}

// Get returns a value of corresponding variable flag.
// Second (bool) parameter says whether it's really defined or not.
func (c *Args) Get(flagName string) (string, bool) {
	s, ok := c.vars[flagName]
	return s, ok
}

// DEPRECATED: Use Has(string) instead.
func (c *Args) Is(flagName string) bool {
	_, ok := c.Get(flagName)
	return ok
}

// Has returns true if a flag with corresponding name is defined.
func (c *Args) Has(flagName string) bool {
	_, ok := c.Get(flagName)
	return ok
}

// String returns a string of corresponding variable flag.
// Second (bool) parameter says whether it's really defined or not.
func (c *Args) String(flagName string) string {
	s, ok := c.Get(flagName)
	if !ok {
		return ""
	}
	return s
}

// Bool returns a bool of corresponding variable flag.
func (c *Args) Bool(flagName string) bool {
	s, ok := c.Get(flagName)
	if !ok {
		return false
	}
	if s == "" {
		return true
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		switch s {
		case "on", "ON", "On":
			return true
		case "off", "OFF", "Off":
			return false
		}
		return false
	}
	return v
}

// Int64 returns a int64 of corresponding variable flag.
//
// Look for prefix of binary("0b"), octal("0o"), hex("0x").
func (c *Args) Int64(flagName string) int64 {
	s, ok := c.Get(flagName)
	if !ok {
		return 0
	}
	n, err := strconv.ParseInt(s, 0, 0)
	if err != nil {
		return 0
	}
	return n
}

// Int returns a int of corresponding variable flag.
//
// Look for prefix of binary("0b"), octal("0o"), hex("0x").
func (c *Args) Int(flagName string) int {
	return int(c.Int64(flagName))
}

// Variables returns a map[string]string of arguments and options of command call.
func (c *Args) Variables() map[string]string {
	return c.vars
}
