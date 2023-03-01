package template

import (
	"io"
	"strings"
)

type LayoutTemplateProcessor struct{}

func (proc *LayoutTemplateProcessor) ExecTemplate(writer io.Writer,
	name string, data interface{}) (err error) {
	return proc.ExecTemplateWithFunc(writer, name, data, emptyFunc)
}

func (proc *LayoutTemplateProcessor) ExecTemplateWithFunc(writer io.Writer,
	name string, data interface{},
	handlerFunc InvokeHandlerFunc) (err error) {

	var sb strings.Builder
	layoutName := ""
	localTemplates := getTemplates()
	localTemplates.Funcs(map[string]interface{}{
		"body":    insertBodyWrapper(&sb),
		"layout":  setLayoutWrapper(&layoutName),
		"handler": handlerFunc,
	})
	err = localTemplates.ExecuteTemplate(&sb, name, data)
	if layoutName != "" {
		localTemplates.ExecuteTemplate(writer, layoutName, data)
	} else {
		io.WriteString(writer, sb.String())
	}
	return
}
