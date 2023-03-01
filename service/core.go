package service

import (
	"context"
	"fmt"
	"reflect"
)

type BindingMap struct {
	factoryFunc reflect.Value
}

var services = make(map[reflect.Type]BindingMap)

func addService(factoryFunc interface{}) (err error) {
	factoryFuncType := reflect.TypeOf(factoryFunc)
	if factoryFuncType.Kind() == reflect.Func && factoryFuncType.NumOut() == 1 {
		services[factoryFuncType.Out(0)] = BindingMap{
			factoryFunc: reflect.ValueOf(factoryFunc),
		}
	} else {
		err = fmt.Errorf("Type cannot be used as service: %v", factoryFuncType)

	}
	return
}

var contextReference = (*context.Context)(nil)
var contextReferenceType = reflect.TypeOf(contextReference).Elem()

func resolveServiceFromValue(c context.Context, val reflect.Value) (err error) {
	serviceType := val.Elem().Type()
	if serviceType == contextReferenceType {
		val.Elem().Set(reflect.ValueOf(c))
	} else if binding, found := services[serviceType]; found {
		val.Elem().Set(invokeFunction(c, binding.factoryFunc)[0])
	} else {
		err = fmt.Errorf("Cannot find service %v", serviceType)
	}
	return
}

func resolveFunctionArguments(c context.Context, f reflect.Value,
	otherArgs ...interface{}) []reflect.Value {
	params := make([]reflect.Value, f.Type().NumIn())
	i := 0
	if otherArgs != nil {
		for ; i < len(otherArgs); i++ {
			params[i] = reflect.ValueOf(otherArgs[i])
		}
	}
	for ; i < len(params); i++ {
		pType := f.Type().In(i)
		pVal := reflect.New(pType)
		err := resolveServiceFromValue(c, pVal)
		if err != nil {
			panic(err)
		}
		params[i] = pVal.Elem()
	}
	return params
}

func invokeFunction(c context.Context, f reflect.Value,
	otherArgs ...interface{}) []reflect.Value {
	return f.Call(resolveFunctionArguments(c, f, otherArgs...))
}
