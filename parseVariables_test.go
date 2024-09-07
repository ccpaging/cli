package cli

import (
	"reflect"
	"testing"
)

func TestContext(t *testing.T) {
	check := func(c string, beStrict bool, f []*Flag, a []string, exp map[string]string) {
		vars, err := parseVariables(beStrict, f, a)
		if err != nil {
			t.Errorf(`case "%s" didn't finish well:`, c)
			t.Logf(`error: %s`, err)
			return
		}

		if !reflect.DeepEqual(vars, exp) {
			t.Errorf(`case "%s" didn't finish well:`, c)
			t.Logf("- expected:\n%v", exp)
			t.Logf("- recieved:\n%v", vars)
		}
	}

	mustFail := func(c string, beStrict bool, f []*Flag, a []string) {
		_, err := parseVariables(beStrict, f, a)
		if err == nil {
			t.Errorf(`invalid case "%s" resulted in valid context`, c)
		}
	}

	// PASS TESTS
	// ==========

	beStrict := false

	check("no args", beStrict, []*Flag{}, []string{}, map[string]string{})

	check("no flags", beStrict, []*Flag{}, []string{"argument", "a thing"}, map[string]string{})

	check("single non-var flag", beStrict, []*Flag{
		{Name: "force"},
	}, []string{"--force", "hard life in ghetto"}, map[string]string{
		"force": "hard life in ghetto",
	})

	check("single shortened non-var flag", beStrict, []*Flag{
		{Name: "force", Short: "f"},
	}, []string{"-f", "hard life in ghetto"}, map[string]string{
		"force": "hard life in ghetto",
	})

	check("separated variable flag", beStrict, []*Flag{
		{Name: "filter"},
	}, []string{"--filter", "token here"}, map[string]string{
		"filter": "token here",
	})

	check("joined empty variable flag", beStrict, []*Flag{
		{Name: "filter"},
	}, []string{`--filter=`}, map[string]string{
		"filter": "",
	})

	check("joined single-word variable flag", beStrict, []*Flag{
		{Name: "filter"},
	}, []string{`--filter=token`}, map[string]string{
		"filter": "token",
	})

	check("joined multi-word variable flag", beStrict, []*Flag{
		{Name: "filter"},
	}, []string{`--filter=token here`}, map[string]string{
		"filter": "token here",
	})

	check("joined single-word shortened variable flag", beStrict, []*Flag{
		{Name: "filter", Short: "f"},
	}, []string{`-f=token`}, map[string]string{
		"filter": "token",
	})

	check("joined multi-word shortened variable flag", beStrict, []*Flag{
		{Name: "filter", Short: "f"},
	}, []string{`-f=token here`}, map[string]string{
		"filter": "token here",
	})

	check("sophisticated", beStrict, []*Flag{
		{Name: "force", Short: "f"},
		{Name: "slug"},
	}, []string{"-f", "--slug", "dog_03", "Dog Doggson"}, map[string]string{
		"force": "",
		"slug":  "dog_03 Dog Doggson",
	})

	check("missing flag", beStrict, []*Flag{
		{Name: "filter"},
	}, []string{"--notexist"}, map[string]string{
		"notexist": "",
	})

	// FAIL TESTS
	// ==========
	beStrict = true

	mustFail("no flags in strict mode", beStrict, []*Flag{
		{Name: "filter"},
	}, []string{"argument", "a thing"})

	mustFail("missing flag in strict mode", beStrict, []*Flag{
		{Name: "filter"},
	}, []string{"--notexist"})
}
