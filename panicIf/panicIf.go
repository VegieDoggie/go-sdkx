package panicIf

import (
	"fmt"
	"github.com/VegetableDoggies/go-sdkx/valx"
	"strings"
)

// Err panics if the arguments include a non-nil err
func Err(arguments ...any) {
	checkErr("", arguments)
}

// Errf panics with format if the arguments include a non-nil err
func Errf(format string, arguments ...any) {
	checkErr(format, arguments)
}

func checkErr(format string, arguments []any) {
loop:
	for i := len(arguments) - 1; i >= 0; i-- {
		switch v := arguments[i].(type) {
		case error:
			if v != nil {
				if len(format) == 0 {
					format = strings.Repeat("%v ", len(arguments))
				}
				panic(fmt.Errorf(format, arguments...))
			}
			break loop
		}
	}
}

// True panics if the arguments include a true bool
func True(arguments ...any) {
	checkBool(true, "", arguments)
}

// Truef panics with format if the arguments include a true bool
func Truef(format string, arguments ...any) {
	checkBool(true, format, arguments)
}

// False panics if the arguments include a false bool
func False(arguments ...any) {
	checkBool(false, "", arguments)
}

// Falsef panics with format if the arguments include a true bool
func Falsef(format string, arguments ...any) {
	checkBool(false, format, arguments)
}

func checkBool(b bool, format string, arguments []any) {
loop:
	for i := len(arguments) - 1; i >= 0; i-- {
		switch v := arguments[i].(type) {
		case bool:
			if v == b {
				if len(format) == 0 {
					format = strings.Repeat("%v ", len(arguments))
				}
				panic(fmt.Errorf(format, arguments...))
			}
			break loop
		}
	}
}

// Nil panics if the first argument is a nil interface, pointer, map, array, slice, chan or func
func Nil(val any, format string, a ...any) {
	if valx.IsNil(val) {
		panic(fmt.Errorf(format, a...))
	}
}

// NotNil panics if the first argument isn't a nil interface, pointer, map, array, slice, chan or func
func NotNil(val any, format string, a ...any) {
	if !valx.IsNil(val) {
		panic(fmt.Errorf(format, a...))
	}
}
