package logIf

import (
	"errors"
	"testing"
)

func TestErr(t *testing.T) {
	err := errors.New("TestErr")
	Err(err, nil)
}

func TestErrf(t *testing.T) {
	err := errors.New("TestErrf")
	Errf("f: %v %v", err, nil)
}

func TestTrue(t *testing.T) {
	err := errors.New("TestTrue")
	True(true, err, nil)
}

func TestTruef(t *testing.T) {
	err := errors.New("TestTruef")
	Truef("f: %v %v %v", true, err, nil)
}

func TestFalse(t *testing.T) {
	err := errors.New("TestFalse")
	False(false, err, nil)
}

func TestFalsef(t *testing.T) {
	err := errors.New("TestFalsef")
	Falsef("f: %v %v %v", false, err, nil)
}

func TestNil(t *testing.T) {
	var m map[string]any
	Nil(m, "f: TestNil")
}

func TestNotNil(t *testing.T) {
	m := make(map[string]any)
	NotNil(m, "f: TestNotNil")
}
