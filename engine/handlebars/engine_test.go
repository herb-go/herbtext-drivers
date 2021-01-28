package handlebars

import (
	"sort"
	"testing"

	"github.com/herb-go/herbtext/texttemplate"

	"github.com/herb-go/herbtext"
)

var testtemplate = `{{test testkey}}-{{prefix 'testprefix' testkey2 }}`

func TestEngine(t *testing.T) {
	eng, err := texttemplate.NewEngine(EngineName, nil)
	if err != nil {
		panic(err)
	}
	env := herbtext.NewEnvironment()
	env.SetConverter("test", func(data string) string {
		return data + "test"
	})

	err = eng.ApplyEnvironment(env)
	if err != nil {
		panic(err)
	}
	supported, err := eng.Supported()
	if err != nil {
		panic(err)
	}
	sort.Strings(supported)
	if len(supported) != 1 || supported[0] != "test" {
		t.Fatal(supported)
	}
	tplenv := herbtext.NewEnvironment()
	tplenv.SetFormatter("prefix", func(prefix string, data string) string {
		return prefix + data
	})
	v, err := eng.Parse(testtemplate, tplenv)
	if err != nil {
		panic(err)
	}
	renderdata := map[string]interface{}{
		"testkey":  "testvalue",
		"testkey2": "testvalue2",
	}
	result, err := v.Render(renderdata)
	if err != nil || result != "testvaluetest-testprefixtestvalue2" {
		t.Fatal(result, err)
	}
}
