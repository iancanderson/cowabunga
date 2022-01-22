package enumerable

import (
	"reflect"
	"strconv"
	"testing"
)

func TestAll(t *testing.T) {
	ints := []int{1, 2, 3}

	cases := []struct {
		in          func(int) bool
		want        bool
		description string
	}{
		{func(n int) bool { return n > 0 }, true, "n > 0"},
		{func(n int) bool { return n > 1 }, false, "n > 1"},
	}
	for _, c := range cases {
		got := All(ints, c.in)
		if got != c.want {
			t.Errorf("All(%v, %v) == %t, want %t", ints, c.description, got, c.want)
		}
	}
}

func TestFilter(t *testing.T) {
	ints := []int{1, 2, 3}

	cases := []struct {
		in          func(int) bool
		want        []int
		description string
	}{
		{func(n int) bool { return n > 0 }, ints, "n > 0"},
		{func(n int) bool { return n > 1 }, []int{2, 3}, "n > 1"},
	}
	for _, c := range cases {
		got := Filter(ints, c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Filter(%v, %v) == %v, want %v", ints, c.description, got, c.want)
		}
	}
}

func TestMap(t *testing.T) {
	ints := []int{1, 2, 3}

	cases := []struct {
		in          func(int) string
		want        []string
		description string
	}{
		{
			func(n int) string { return strconv.Itoa(n + 1) },
			[]string{"2", "3", "4"},
			"n + 1",
		},
	}
	for _, c := range cases {
		got := Map(ints, c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Map(%v, %v) == %v, want %v", ints, c.description, got, c.want)
		}
	}
}
