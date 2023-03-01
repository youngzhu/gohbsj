package service

import "youngzy.com/gohbsj/template"

func NewTemplateAction(name string, data interface{}) ActionResult {
	return &TemplateActionResult{templateName: name, data: data}
}

type TemplateActionResult struct {
	templateName string
	data         interface{}
	template.TemplateExecutor
	template.InvokeHandlerFunc
}

func (action *TemplateActionResult) Execute(ctx *ActionContext) error {
	return action.TemplateExecutor.ExecTemplateWithFunc(ctx.ResponseWriter,
		action.templateName, action.data, action.InvokeHandlerFunc)
}
