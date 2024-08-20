package logIf

import (
	"errors"
	"testing"
)

func TestErr(t *testing.T) {
	Err(1, 2, 3, errors.New("test"))
}

func TestTrue(t *testing.T) {
	True(true, "test")
}

func TestFalse(t *testing.T) {
	False(false, "test")
}

func TestNil(t *testing.T) {
	Nil(false, "test")
}

func TestEmpty(t *testing.T) {
	Empty(make(map[string]any), "test")
}
