package panicIf

import (
	"fmt"
	"github.com/VegieDoggie/go-sdkx/check"
	"reflect"
)

var validationFailedError = fmt.Errorf("validation failed! ")

// Err log if the last argument is a non-nil error
func Err(err any, msg ...any) {
	if err != nil {
		switch err.(type) {
		case error:
			panicErr("Err", err.(error), msg)
		}
	}
}

// ErrReturn panic if the last argument is a non-nil error
func ErrReturn(arguments ...any) {
	if len(arguments) > 0 {
		if err := arguments[len(arguments)-1]; err != nil {
			switch err.(type) {
			case error:
				panicErr("ErrReturn", err.(error), nil)
			}
		}
	}
}

// ErrNil panic if err error or v is nil
func ErrNil(err any, obj any, msg ...any) {
	Err(err, msg...)
	Nil(obj, msg...)
}

// True panic if the argument is true
func True(argument bool, msg ...any) {
	if argument {
		panicErr("True", validationFailedError, msg)
	}
}

// False panic if the argument is false
func False(argument bool, msg ...any) {
	if !argument {
		panicErr("False", validationFailedError, msg)
	}
}

// Nil panic if the argument is nil
func Nil(argument any, msg ...any) {
	if argument == nil {
		panicErr("Nil", validationFailedError, msg)
		return
	}

	val := reflect.ValueOf(argument)
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if val.IsNil() {
			panicErr("Nil", validationFailedError, msg)
		}
	default:
	}
}

// Empty panic if the argument is nil or empty
func Empty(argument any, msg ...any) {
	if check.IsEmpty(argument) {
		panicErr("Empty", validationFailedError, msg)
	}
}

func panicErr(typ string, rawErr error, errMsg []any) {
	if len(errMsg) > 0 {
		switch errMsg[0].(type) {
		case error:
			panic(errMsg[0]) // custom error
		}
		panic(fmt.Errorf("panicErr %v: %v\n%+v", typ, rawErr, errMsg))
	}
	panic(fmt.Errorf("panicErr %v: %v", typ, rawErr))
}
