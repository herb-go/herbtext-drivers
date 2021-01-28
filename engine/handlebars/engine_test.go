package handlebars

import (
	"sort"
	"testing"

	"github.com/herb-go/herbtext/texttemplate"

	"github.com/herb-go/herbtext"
)

var testtemplate = `{{test testkey}}-{{prefix 'testprefix' testkey2 }}`

func TestEngine(t *testing.T) {
	eng, err := texttemplate.GetEngine(EngineName)
	if err != nil {
		panic(err)
	}
	env := herbtext.NewEnvironment()
	env.SetConverter("test", func(data string) string {
		return data + "test"
	})
	env.SetFormatter("prefix", func(prefix string, data string) string {
		return prefix + data
	})
	supported, err := eng.Supported(env)
	if err != nil {
		panic(err)
	}
	sort.Strings(supported)
	if len(supported) != 2 || supported[0] != "prefix" || supported[1] != "test" {
		t.Fatal(supported)
	}

	v, err := eng.Parse(testtemplate, env)
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
