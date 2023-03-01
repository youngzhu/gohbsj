package service

import (
	"context"
	"errors"
	"reflect"
)

func Call(target interface{}, otherArgs ...interface{}) ([]interface{}, error) {
	return CallForContext(context.Background(), target, otherArgs...)
}

func CallForContext(c context.Context, target interface{}, otherArgs ...interface{}) (results []interface{}, err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Func {
		resultVals := invokeFunction(c, targetValue, otherArgs...)
		results = make([]interface{}, len(resultVals))
		for i := 0; i < len(resultVals); i++ {
			results[i] = resultVals[i].Interface()
		}
	} else {
		err = errors.New("Only functions can be invoked")
	}
	return
}

func GetService(target interface{}) error {
	return GetServiceForContext(context.Background(), target)
}

func GetServiceForContext(c context.Context, target interface{}) (err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr &&
		targetValue.Elem().CanSet() {
		err = resolveServiceFromValue(c, targetValue)
	} else {
		err = errors.New("Type cannot be used as target")
	}
	return
}
