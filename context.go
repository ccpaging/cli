package cli

import (
	"strconv"
)

// Context is a set of arguments and options of command call.
type Context struct {
	app  *App
	vars map[string]string
}

func newContext(a *App, flags []*Flag, argv []string) (*Context, error) {
	vars, err := parseVariables(a.Strict, flags, argv)
	if err != nil {
		return nil, err
	}

	c := &Context{
		app:  a,
		vars: vars,
	}
	return c, nil
}

// Get returns a value of corresponding variable flag.
// Second (bool) parameter says whether it's really defined or not.
func (c *Context) Get(flagName string) (string, bool) {
	s, ok := c.vars[flagName]
	return s, ok
}

// DEPRECATED: Use Has(string) instead.
func (c *Context) Is(flagName string) bool {
	_, ok := c.Get(flagName)
	return ok
}

// Has returns true if a flag with corresponding name is defined.
func (c *Context) Has(flagName string) bool {
	_, ok := c.Get(flagName)
	return ok
}

// String returns a string of corresponding variable flag.
// Second (bool) parameter says whether it's really defined or not.
func (c *Context) String(flagName string) string {
	s, ok := c.Get(flagName)
	if !ok {
		return ""
	}
	return s
}

// Bool returns a bool of corresponding variable flag.
func (c *Context) Bool(flagName string) bool {
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
func (c *Context) Int64(flagName string) int64 {
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
func (c *Context) Int(flagName string) int {
	return int(c.Int64(flagName))
}

// MapString returns a map[string]string of arguments and options of command call.
func (c *Context) MapString() map[string]string {
	return c.vars
}

// MapInterface returns a map[string]interface{} of arguments and options of command call.
func (c *Context) MapInterface() map[string]interface{} {
	out := make(map[string]interface{})
	for k, v := range c.vars {
		out[k] = v
	}
	return out
}
