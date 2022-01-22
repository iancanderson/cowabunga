package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Returns true if all elements in input
// give a truthy result when passed to f
func All[T any](input []T, f func(T) bool) bool {
	for _, el := range input {
		if !f(el) {
			return false
		}
	}
	return true
}

func Map[T, K any](input []T, f func(T) K) []K {
	output := make([]K, len(input))
	for i, el := range input {
		output[i] = f(el)
	}
	return output
}

func Filter[T any](input []T, f func(T) bool) []T {
	output := make([]T, 1)
	for _, el := range input {
		if f(el) {
			output = append(output, el)
		}
	}
	return output
}

func main() {
	strs := []string{"foo", "bar", "baz"}
	ints := []int{1, 2, 3}

	fmt.Println(All(strs, func(s string) bool {
		return len(s) == 3
	}))
	fmt.Println(All(strs, func(s string) bool {
		return s[0] == 'f'
	}))

	fmt.Println(Map(strs, func(s string) string {
		return strings.ToUpper(s)
	}))
	fmt.Println(Map(ints, func(i int) string {
		return strconv.Itoa(i + 1)
	}))

	fmt.Println(Filter(strs, func(s string) bool {
		return s[0] == 'b'
	}))
}
