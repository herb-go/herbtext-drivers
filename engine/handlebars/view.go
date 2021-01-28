package handlebars

import (
	"github.com/aymerick/raymond"
)

//View handlebars template view
type View struct {
	template  *raymond.Template
	supported []string
}

//Render render given data and return output string and any error if raised.
func (v *View) Render(data interface{}) (output string, err error) {
	return v.template.Exec(data)
}
