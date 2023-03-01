package basic

import (
	"fmt"
	"net/http"
	"youngzy.com/gohbsj/logging"
	"youngzy.com/gohbsj/service"
)

type ErrorComponent struct{}

func recoveryFunc(ctx *service.ComponentContext, logger logging.Logger) {
	if arg := recover(); arg != nil {
		logger.Debugf("Error: %v", fmt.Sprint(arg))
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *ErrorComponent) Init() {}

func (c *ErrorComponent) ProcessRequest(ctx *service.ComponentContext,
	next func(*service.ComponentContext)) {

	var logger logging.Logger
	service.GetServiceForContext(ctx.Context(), &logger)
	//defer recoveryFunc(ctx, logger)
	next(ctx)
	if ctx.GetError() != nil {
		logger.Debugf("Error: %v", ctx.GetError())
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
