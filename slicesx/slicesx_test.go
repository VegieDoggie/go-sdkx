package slicesx

import (
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		s    []int
		f    func(int) bool
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			func(i int) bool {
				return i < 3
			},
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3, 4, 5},
			func(i int) bool {
				return i > 3
			},
			[]int{4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			func(i int) bool {
				return i == 3
			},
			[]int{3},
		},
	}
	for i := range tests {
		if got := Filter(tests[i].s, tests[i].f); !Equal(got, tests[i].want) {
			t.Errorf("Filter(%v, tests[i].f) = %v, want = %v", tests[i].s, got, tests[i].want)
		}
	}
}

func TestMap(t *testing.T) {
	type Xxx struct {
		A string
		B int
	}
	tests := []Xxx{
		{
			"1",
			10,
		},
		{
			"2",
			20,
		},
		{
			"3",
			30,
		},
	}
	t1 := Map(tests, func(t Xxx) string {
		return t.A
	})
	w1 := []string{"1", "2", "3"}
	if !Equal(t1, w1) {
		t.Errorf("Map(%v, func(t Xxx) string) = %v, want = %v", tests, t1, w1)
	}
	t2 := Map(tests, func(t Xxx) int {
		return t.B
	})
	w2 := []int{10, 20, 30}
	if !Equal(t1, w1) {
		t.Errorf("Map(%v, func(t Xxx) string) = %v, want = %v", tests, t2, w2)
	}
}

func Equal[E comparable](s1, s2 []E) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
