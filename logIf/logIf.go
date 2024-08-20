package logIf

import (
	"github.com/VegieDoggie/go-sdkx/check"
	"log"
	"os"
	"reflect"
)

var std = log.New(os.Stderr, "<logIf> ", log.Ldate|log.Ltime)

func SetLogger(logger *log.Logger) {
	std = logger
}

// Err log if the last argument is a non-nil error
func Err(argument any, info ...string) {
	if argument != nil {
		switch argument.(type) {
		case error:
			std.Printf("log Err: %+v\n%v\n", info, argument)
		}
	}
}

// ErrReturned log if the last argument is a non-nil error
func ErrReturned(arguments ...any) {
	if len(arguments) > 0 {
		if a := arguments[len(arguments)-1]; a != nil {
			switch a.(type) {
			case error:
				std.Printf("log ErrReturned: %+v\n", a)
			}
		}
	}
}

// ErrNil log if err error or v is nil
func ErrNil(err any, vNil any, info ...string) {
	Err(err, info...)
	Nil(vNil, info...)
}

// True log if the argument is true
func True(argument bool, info ...string) {
	if argument {
		std.Printf("log True: %+v\n", info)
	}
}

// False log if the argument is false
func False(argument bool, info ...string) {
	if !argument {
		std.Printf("log False: %+v\n", info)
	}
}

// Nil log if the argument is nil
func Nil(argument any, info ...string) {
	if argument == nil || reflect.ValueOf(argument).IsNil() {
		std.Printf("log Nil: %+v\n", info)
	}
}

// Empty log if the argument is nil or empty
func Empty(argument any, info ...string) {
	if check.IsEmpty(argument) {
		std.Printf("log Empty: %+v\n", info)
	}
}
