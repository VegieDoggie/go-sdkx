package panicIf

import (
	"fmt"
	"github.com/VegetableDoggies/go-sdkx/sdkx"
	"strings"
)

// Err panics if the last argument is a non-nil error
func Err(arguments ...any) {
	if a := arguments[len(arguments)-1]; a != nil {
		switch a.(type) {
		case error:
			panic(a)
		}
	}
}

// Errf panics with format if the last argument is a non-nil error
func Errf(format string, arguments ...any) {
	if a := arguments[len(arguments)-1]; a != nil {
		switch a.(type) {
		case error:
			panic(fmt.Errorf(format, arguments...))
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

// Nil logs if the last argument is nil
func Nil(format string, arguments ...any) {
	if arguments[len(arguments)-1] == nil {
		panic(fmt.Errorf(format, arguments...))
	}
}

// NotNil logs if the last argument isn't nil
func NotNil(format string, arguments ...any) {
	if arguments[len(arguments)-1] != nil {
		panic(fmt.Errorf(format, arguments...))
	}
}

// Empty panics if the last argument is a nil interface, pointer, or empty map, array, slice, chan, string
func Empty(format string, arguments ...any) {
	if sdkx.IsEmpty(arguments[len(arguments)-1]) {
		panic(fmt.Errorf(format, arguments...))
	}
}

// NotEmpty panics if the last argument isn't a nil interface, pointer, or empty map, array, slice, chan, string
func NotEmpty(format string, arguments ...any) {
	if !sdkx.IsEmpty(arguments[len(arguments)-1]) {
		panic(fmt.Errorf(format, arguments...))
	}
}
