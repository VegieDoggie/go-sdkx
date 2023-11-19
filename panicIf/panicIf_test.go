package panicIf

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func assertPanic(s string) {
	if r := recover(); r != nil {
		if !strings.Contains(r.(error).Error(), s) {
			panic(fmt.Sprintf("unexpected: %v not in [%v]", s, r))
		} else {
			fmt.Println(fmt.Sprintf("assertPanic: %v in [%v]", s, r))
		}
	}
}

func TestErr(t *testing.T) {
	defer assertPanic("TestErr")
	err := errors.New("TestErr")
	Err(err, nil)
}

func TestErrf(t *testing.T) {
	defer assertPanic("TestErrf")
	err := errors.New("TestErrf")
	Errf("f: %v %v", err, nil)
}

func TestTrue(t *testing.T) {
	defer assertPanic("true")
	err := errors.New("TestTrue")
	True(true, err, nil)
}

func TestTruef(t *testing.T) {
	defer assertPanic("true")
	err := errors.New("TestTruef")
	Truef("f: %v %v %v", true, err, nil)
}

func TestFalse(t *testing.T) {
	defer assertPanic("false")
	err := errors.New("TestFalse")
	False(false, err, nil)
}

func TestFalsef(t *testing.T) {
	defer assertPanic("false")
	err := errors.New("TestFalsef")
	Falsef("f: %v %v %v", false, err, nil)
}

func TestNil(t *testing.T) {
	defer assertPanic("TestNil")
	var m map[string]any
	Nil("f: TestNil", m)
}

func TestNotNil(t *testing.T) {
	defer assertPanic("TestNotNil")
	m := make(map[string]any)
	NotNil("f: TestNotNil", m)
}
