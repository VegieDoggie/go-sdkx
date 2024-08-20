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
	defer assertPanic("test")
	Err(errors.New("test"))
}

func TestTrue(t *testing.T) {
	defer assertPanic("test")
	True(true, "test")
}

func TestFalse(t *testing.T) {
	defer assertPanic("test")
	False(false, "test")
}

func TestNil(t *testing.T) {
	defer assertPanic("test")
	Nil(false, "test")
}

func TestEmpty(t *testing.T) {
	defer assertPanic("test")
	Empty(make(map[string]any), "test")
}
