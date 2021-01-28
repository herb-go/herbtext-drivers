package handlebars

import (
	"github.com/aymerick/raymond"
	"github.com/herb-go/herbtext"
	"github.com/herb-go/herbtext/texttemplate"
)

//EngineName engine name
const EngineName = "handlebars"

func applyEnvironment(tpl *raymond.Template, env herbtext.Environment) {
	env.RangeConverters(func(key string, converter herbtext.Converter) bool {
		tpl.RegisterHelper(key, converter)
		return true
	})
	env.RangeFormatters(func(key string, formatter herbtext.Formatter) bool {
		tpl.RegisterHelper(key, formatter)
		return true
	})

}
func supported(env herbtext.Environment) []string {
	var result []string
	env.RangeConverters(func(key string, converter herbtext.Converter) bool {
		result = append(result, key)
		return true
	})
	env.RangeFormatters(func(key string, formatter herbtext.Formatter) bool {
		result = append(result, key)
		return true
	})
	return result
}

//Engine engine struct
type Engine struct {
	env herbtext.Environment
}

//ApplyEnvironment apply given environment to engine and return any error if raised.
func (e *Engine) ApplyEnvironment(env herbtext.Environment) error {
	e.env = herbtext.Clone(env)
	return nil
}

//Parse parse given template with given environment to template view.
func (e *Engine) Parse(template string, env herbtext.Environment) (texttemplate.View, error) {
	tpl, err := raymond.Parse(template)
	if err != nil {
		return nil, err
	}
	tplenv := herbtext.Clone(e.env)
	tplenv.MergeWith(env)
	applyEnvironment(tpl, tplenv)
	return &View{template: tpl, supported: supported(tplenv)}, nil
}

//Supported return supported directives which can be used in template string.
func (e *Engine) Supported() (directives []string, err error) {
	return supported(e.env), nil
}

//Factory engine factory function.
func Factory(loader func(v interface{}) error) (texttemplate.Engine, error) {
	return &Engine{
		env: herbtext.DefaultEnvironment(),
	}, nil
}

func init() {
	texttemplate.Register(EngineName, Factory)
}
