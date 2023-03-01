package basic

import (
	"net/http"
	"youngzy.com/gohbsj/service"

	//"platform/services"
	"strings"
)

type StaticFileComponent struct {
	urlPrefix     string
	stdLibHandler http.Handler
}

func (sfc *StaticFileComponent) Init() {
	sfc.urlPrefix = "/resources/"
	sfc.stdLibHandler = http.StripPrefix(sfc.urlPrefix,
		http.FileServer(http.Dir("./resources")))
}

func (sfc *StaticFileComponent) ProcessRequest(ctx *service.ComponentContext,
	next func(*service.ComponentContext)) {

	if !strings.EqualFold(ctx.Request.URL.Path, sfc.urlPrefix) &&
		strings.HasPrefix(ctx.Request.URL.Path, sfc.urlPrefix) {
		sfc.stdLibHandler.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	} else {
		next(ctx)
	}
}
