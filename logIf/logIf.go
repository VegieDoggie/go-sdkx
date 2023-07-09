package logIf

import (
	"fmt"
	"github.com/VegetableDoggies/go-sdkx/sdkx"
	"log"
	"os"
	"strings"
)

var std = log.New(os.Stderr, "<logIf> ", log.Ldate|log.Ltime)

func SetLogger(logger *log.Logger) {
	std = logger
}

// Err logs if the arguments include a non-nil err
func Err(arguments ...any) {
	checkErr("", arguments)
}

// Errf logs with format if the arguments include a non-nil err
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
				std.Println(fmt.Sprintf(format, arguments...))
			}
			break loop
		}
	}
}

// True logs if the arguments include a true bool
func True(arguments ...any) {
	checkBool(true, "", arguments)
}

// Truef logs with format if the arguments include a true bool
func Truef(format string, arguments ...any) {
	checkBool(true, format, arguments)
}

// False logs if the arguments include a false bool
func False(arguments ...any) {
	checkBool(false, "", arguments)
}

// Falsef logs with format if the arguments include a true bool
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
				std.Println(fmt.Sprintf(format, arguments...))
			}
			break loop
		}
	}
}

// Nil logs if the last argument is nil
func Nil(format string, arguments ...any) {
	if arguments[len(arguments)-1] == nil {
		std.Println(fmt.Sprintf(format, arguments[:len(arguments)-1]...))
	}
}

// NotNil logs if the last argument isn't nil
func NotNil(format string, arguments ...any) {
	if arguments[len(arguments)-1] != nil {
		std.Println(fmt.Sprintf(format, arguments...))
	}
}

// Empty panics if the last argument is a nil interface, pointer, or empty map, array, slice, chan, string
func Empty(format string, arguments ...any) {
	if sdkx.IsEmpty(arguments[len(arguments)-1]) {
		std.Println(fmt.Sprintf(format, arguments[:len(arguments)-1]...))
	}
}

// NotEmpty panics if the last argument isn't a nil interface, pointer, or empty map, array, slice, chan, string
func NotEmpty(format string, arguments ...any) {
	if !sdkx.IsEmpty(arguments[len(arguments)-1]) {
		std.Println(fmt.Sprintf(format, arguments...))
	}
}
