package service

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"youngzy.com/gohbsj/service/param"
)

func NewRouter(handlers ...HandlerEntry) *RouterComponent {
	routes := generateRoutes(handlers...)

	var urlGen URLGenerator
	GetService(&urlGen)
	if urlGen == nil {
		AddSingleton(func() URLGenerator {
			return &routeUrlGenerator{routes: routes}
		})
	} else {
		urlGen.AddRoutes(routes)
	}
	return &RouterComponent{routes: routes}
}

type RouterComponent struct {
	routes []Route
}

func (router *RouterComponent) Init() {}

func (router *RouterComponent) ProcessRequest(context *ComponentContext,
	next func(*ComponentContext)) {
	for _, route := range router.routes {
		if strings.EqualFold(context.Request.Method, route.httpMethod) {
			matches := route.expression.FindAllStringSubmatch(context.URL.Path, -1)
			if len(matches) > 0 {
				rawParamVals := []string{}
				if len(matches[0]) > 1 {
					rawParamVals = matches[0][1:]
				}
				err := router.invokeHandler(route, rawParamVals, context)
				if err == nil {
					next(context)
				} else {
					context.Error(err)
				}
				return
			}
		}
	}
	context.ResponseWriter.WriteHeader(http.StatusNotFound)
}

func (router *RouterComponent) invokeHandler(route Route, rawParams []string,
	context *ComponentContext) error {
	paramVals, err := param.GetParametersFromRequest(context.Request,
		route.handlerMethod, rawParams)
	if err == nil {
		structVal := reflect.New(route.handlerMethod.Type.In(0))
		PopulateForContext(context.Context(), structVal.Interface())
		paramVals = append([]reflect.Value{structVal.Elem()}, paramVals...)
		result := route.handlerMethod.Func.Call(paramVals)
		if len(result) > 0 {
			if action, ok := result[0].Interface().(ActionResult); ok {
				invoker := createInvokehandlerFunc(context.Context(), router.routes)
				err = PopulateForContextWithExtras(context.Context(),
					action,
					map[reflect.Type]reflect.Value{
						reflect.TypeOf(invoker): reflect.ValueOf(invoker),
					})
				if err == nil {
					err = action.Execute(&ActionContext{
						context.Context(), context.ResponseWriter})
				}
			} else {
				io.WriteString(context.ResponseWriter,
					fmt.Sprint(result[0].Interface()))
			}
		}
	}
	return err
}
