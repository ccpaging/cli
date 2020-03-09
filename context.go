package cli

import (
	"errors"
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

// Has returns true if a flag with corresponding name is defined.
func (c *Context) Has(flagName string) bool {
	if _, ok := c.vars[flagName]; ok {
		return true
	}

	return false
}

// DEPRECATED: Use String(string) instead.
func (c *Context) ValueOf(flagName string) (string, bool) {
	return c.String(flagName)
}

// String returns a string of corresponding variable flag.
// Second (bool) parameter says whether it's really defined or not.
func (c *Context) String(flagName string) (string, bool) {
	s, ok := c.vars[flagName]
	return s, ok
}

// Bool returns a bool of corresponding variable flag.
func (c *Context) Bool(flagName string) bool {
	s, ok := c.vars[flagName]
	if !ok {
		return false
	}
	ok, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return ok
}

// Int64 returns a int64 of corresponding variable flag.
func (c *Context) Int64(flagName string) (int64, error) {
	s, ok := c.vars[flagName]
	if !ok {
		return 0, errors.New("Not found " + flagName)
	}
	return strconv.ParseInt(s, 0, 0)
}

// Int returns a int of corresponding variable flag.
func (c *Context) Int(flagName string) (int, error) {
	i64, err := c.Int64(flagName)
	return int(i64), err
}

// Vars returns a map[string]string of arguments and options of command call.
func (c *Context) Vars() map[string]string {
	return c.vars
}

// Map returns a map[string]interface{} of arguments and options of command call.
func (c *Context) Map() map[string]interface{} {
	out := make(map[string]interface{})
	for k, v := range c.vars {
		out[k] = v
	}
	return out
}
