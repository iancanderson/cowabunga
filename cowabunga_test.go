package cowabunga

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

func TestAny(t *testing.T) {
	ints := []int{1, 2, 3}

	cases := []struct {
		in          func(int) bool
		want        bool
		description string
	}{
		{func(n int) bool { return n > 0 }, true, "n > 0"},
		{func(n int) bool { return n > 1 }, true, "n > 1"},
		{func(n int) bool { return n > 3 }, false, "n > 3"},
	}
	for _, c := range cases {
		got := Any(ints, c.in)
		if got != c.want {
			t.Errorf("Any(%v, %v) == %t, want %t", ints, c.description, got, c.want)
		}
	}
}

func TestCount(t *testing.T) {
	ints := []int{1, 2, 3}

	cases := []struct {
		in          func(int) bool
		want        int
		description string
	}{
		{func(n int) bool { return n > 0 }, 3, "n > 0"},
		{func(n int) bool { return n > 1 }, 2, "n > 1"},
		{func(n int) bool { return n > 3 }, 0, "n > 3"},
	}
	for _, c := range cases {
		got := Count(ints, c.in)
		if got != c.want {
			t.Errorf("Count(%v, %v) == %v, want %v", ints, c.description, got, c.want)
		}
	}
}

func TestDrop(t *testing.T) {
	ints := []int{1, 2, 3}

	cases := []struct {
		in   int
		want []int
	}{
		{0, []int{1, 2, 3}},
		{1, []int{2, 3}},
		{2, []int{3}},
	}
	for _, c := range cases {
		got := Drop(ints, c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Drop(%v, %v) == %v, want %v", ints, c.in, got, c.want)
		}
	}
}

func TestDropWhile(t *testing.T) {
	ints := []int{1, 2, 3}

	cases := []struct {
		in          func(int) bool
		want        []int
		description string
	}{
		{func(n int) bool { return n < 1 }, ints, "n < 1"},
		{func(n int) bool { return n < 3 }, []int{3}, "n < 3"},
	}
	for _, c := range cases {
		got := DropWhile(ints, c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("DropWhile(%v, %v) == %v, want %v", ints, c.description, got, c.want)
		}
	}
}

func TestEachCons(t *testing.T) {
	ints := []int{1, 2, 3}

	got := [][]int{}
	EachCons(ints, 2, func(sequence []int) {
		got = append(got, sequence)
	})

	want := [][]int{{1, 2}, {2, 3}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("EachCons(%v, %v) == %v, want %v", ints, 2, got, want)
	}
}

func TestEachSlice(t *testing.T) {
	ints := []int{1, 2, 3}

	got := [][]int{}
	EachSlice(ints, 2, func(sequence []int) {
		got = append(got, sequence)
	})

	want := [][]int{{1, 2}, {3}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("EachSlice(%v, %v) == %v, want %v", ints, 2, got, want)
	}
}

func TestEachWith(t *testing.T) {
	ints := []int{1, 2, 3}
	with := 0

	got := EachWith(ints, &with, func(n int, result *int) {
		*result += n
	})

	want := 6
	if *got != want {
		t.Errorf("EachWith(%v, %v) == %v, want %v", ints, with, got, want)
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

func TestFilterMap(t *testing.T) {
	ints := []int{1, 2, 3}

	cases := []struct {
		in          func(int) *string
		want        []string
		description string
	}{
		{
			func(n int) *string {
				if n > 1 {
					str := strconv.Itoa(n + 1)
					return &str
				} else {
					return nil
				}
			},
			[]string{"3", "4"},
			"n + 1 if n > 1",
		},
	}
	for _, c := range cases {
		got := FilterMap(ints, c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FilterMap(%v, %v) == %v, want %v", ints, c.description, got, c.want)
		}
	}
}

func TestFind(t *testing.T) {
	ints := []int{1, 2, 3}
	two := 2

	cases := []struct {
		in          func(int) bool
		want        *int
		description string
	}{
		{
			func(n int) bool { return n > 1 },
			&two,
			"n > 1",
		},
		{
			func(n int) bool { return n > 3 },
			nil,
			"n > 3",
		},
	}
	for _, c := range cases {
		got := Find(ints, c.in)
		if !(got == nil && c.want == nil) && *got != *c.want {
			t.Errorf("Find(%v, %v) == %v, want %v", ints, c.description, got, c.want)
		}
	}
}

func TestFirst(t *testing.T) {
	one := 1
	cases := []struct {
		in   []int
		want *int
	}{
		{[]int{1, 2, 3}, &one},
		{[]int{}, nil},
	}
	for _, c := range cases {
		got := First(c.in)
		if !(got == nil && c.want == nil) && *got != *c.want {
			t.Errorf("First(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestIsMember(t *testing.T) {
	ints := []int{1, 2, 3}

	cases := []struct {
		in   int
		want bool
	}{
		{1, true},
		{4, false},
	}
	for _, c := range cases {
		got := IsMember(ints, c.in)
		if got != c.want {
			t.Errorf("IsMember(%v, %v) == %t, want %t", ints, c.in, got, c.want)
		}
	}
}

func TestLast(t *testing.T) {
	three := 3
	cases := []struct {
		in   []int
		want *int
	}{
		{[]int{1, 2, 3}, &three},
		{[]int{}, nil},
	}
	for _, c := range cases {
		got := Last(c.in)
		if !(got == nil && c.want == nil) && *got != *c.want {
			t.Errorf("Last(%v) == %v, want %v", c.in, got, c.want)
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

func TestTake(t *testing.T) {
	ints := []int{1, 2, 3}

	cases := []struct {
		in   int
		want []int
	}{
		{0, []int{}},
		{1, []int{1}},
		{2, []int{1, 2}},
	}
	for _, c := range cases {
		got := Take(ints, c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Take(%v, %v) == %v, want %v", ints, c.in, got, c.want)
		}
	}
}
