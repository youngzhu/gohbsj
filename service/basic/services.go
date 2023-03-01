package basic

import (
	"youngzy.com/gohbsj/service"
)

type ServicesComponent struct{}

func (c *ServicesComponent) Init() {}

func (c *ServicesComponent) ProcessRequest(ctx *service.ComponentContext,
	next func(*service.ComponentContext)) {
	reqContext := ctx.Request.Context()
	ctx.Request.WithContext(service.NewServiceContext(reqContext))
	next(ctx)
}
