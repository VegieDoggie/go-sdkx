package panicIf

import (
	"fmt"
	"github.com/VegieDoggie/go-sdkx/check"
	"reflect"
)

// Err log if the last argument is a non-nil error
func Err(err any, msg ...any) {
	if err != nil {
		switch err.(type) {
		case error:
			panicErr("Err", msg)
		}
	}
}

// ErrReturn panic if the last argument is a non-nil error
func ErrReturn(arguments ...any) {
	if len(arguments) > 0 {
		if err := arguments[len(arguments)-1]; err != nil {
			switch err.(type) {
			case error:
				panicErr("ErrReturn", nil)
			}
		}
	}
}

// ErrNil panic if err error or v is nil
func ErrNil(err any, nl any, msg ...any) {
	Err(err, msg...)
	Nil(nl, msg...)
}

// True panic if the argument is true
func True(argument bool, msg ...any) {
	if argument {
		panicErr("True", msg)
	}
}

// False panic if the argument is false
func False(argument bool, msg ...any) {
	if !argument {
		panicErr("False", msg)
	}
}

// Nil panic if the argument is nil
func Nil(argument any, msg ...any) {
	if argument == nil || reflect.ValueOf(argument).IsNil() {
		panicErr("Nil", msg)
	}
}

// Empty panic if the argument is nil or empty
func Empty(argument any, msg ...any) {
	if check.IsEmpty(argument) {
		panicErr("Empty", msg)
	}
}

func panicErr(typ string, errMsg []any) {
	if len(errMsg) > 0 {
		switch errMsg[0].(type) {
		case error:
			panic(errMsg[0]) // custom error
		}
		panic(fmt.Errorf("panicIf %v: %+v", typ, errMsg))
	}
	panic(fmt.Errorf("panicIf %v", typ))
}
