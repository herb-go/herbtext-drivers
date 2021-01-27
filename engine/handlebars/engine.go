package handlebars

import (
	"github.com/aymerick/raymond"
	"github.com/herb-go/herbtext"
	"github.com/herb-go/herbtext/texttemplate"
)

type Engine struct {
	raymond.Options
}

func (e *Engine) ApplyOptions(*herbtext.Environment) error {

}
func (e *Engine) Parse(template string, env *herbtext.Environment) (texttemplate.View, error) {

}
func (e *Engine) Supported() (directives []string, err error) {

}
