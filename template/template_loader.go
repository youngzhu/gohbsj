package template

import (
	"html/template"
	"sync"
)

const templatePath = "resources/template/*.html"

var once = sync.Once{}

func LoadTemplates() (err error) {
	once.Do(func() {
		parseTemplates := func() (t *template.Template) {
			t = template.New("htmlTemplates")
			t.Funcs(map[string]interface{}{
				"layout":  func() string { return "" },
				"body":    func() string { return "" },
				"handler": func() interface{} { return "" },
			})
			t, err = t.ParseGlob(templatePath)
			return
		}

		templates := parseTemplates()
		getTemplates = func() *template.Template {
			t, _ := templates.Clone()
			return t
		}
		return
	})

	return
}
