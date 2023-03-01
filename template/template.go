package template

import (
	"html/template"
	"io"
	"strings"
)

type TemplateExecutor interface {
	ExecTemplate(writer io.Writer, name string, data interface{}) (err error)

	ExecTemplateWithFunc(writer io.Writer, name string,
		data interface{}, handlerFunc InvokeHandlerFunc) (err error)
}

type InvokeHandlerFunc func(handlerName string, methodName string,
	args ...interface{}) interface{}

var emptyFunc = func(handlerName, methodName string,
	args ...interface{}) interface{} {
	return ""
}

var getTemplates func() (t *template.Template)

func insertBodyWrapper(body *strings.Builder) func() template.HTML {
	return func() template.HTML {
		return template.HTML(body.String())
	}
}

func setLayoutWrapper(val *string) func(string) string {
	return func(layout string) string {
		*val = layout
		return ""
	}
}
