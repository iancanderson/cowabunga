package cowabunga

import (
	"reflect"
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

// Returns true if any elements in input
// give a truthy result when passed to f
func Any[T any](input []T, f func(T) bool) bool {
	for _, el := range input {
		if f(el) {
			return true
		}
	}
	return false
}

func Count[T any](input []T, f func(T) bool) int {
	result := 0
	for _, el := range input {
		if f(el) {
			result++
		}
	}
	return result
}

// Return a new slice excluding the first n elements
// of input
func Drop[T any](input []T, n int) []T {
	return input[n:]
}

// Return a new slice excluding the first n elements
// of input where f is true
func DropWhile[T any](input []T, f func(T) bool) []T {
	n := 0
	for _, el := range input {
		if f(el) {
			n++
		}
	}
	return input[n:]
}

// Calls the function with each successive overlapped n-tuple
// of elements
func EachCons[T any](input []T, n int, f func([]T)) {
	inputLen := len(input)
	for start := range input {
		if start+n > inputLen {
			return
		} else {
			f(input[start : start+n])
		}
	}
}

// Calls the function with each successive disjoint n-tuple
// of elements
func EachSlice[T any](input []T, n int, f func([]T)) {
	inputLen := len(input)
	for start := 0; start < inputLen; start += n {
		if start+n < inputLen {
			f(input[start : start+n])
		} else {
			f(input[start:inputLen])
		}
	}
}

// Calls the function once for each element,
// passing both the element and the given argument
func EachWith[T, W any](input []T, with W, f func(T, W)) W {
	for _, el := range input {
		f(el, with)
	}
	return with
}

// Returns a new slice containing all elements of input
// that give a truthy result when passed to function f
func Filter[T any](input []T, f func(T) bool) []T {
	output := make([]T, 0)
	for _, el := range input {
		if f(el) {
			output = append(output, el)
		}
	}
	return output
}

// Calls the function with successive elements;
// returns an array containing each truthy value returned by the block
func FilterMap[T, K any](input []T, f func(T) *K) []K {
	output := make([]K, 0, len(input))
	for _, el := range input {
		val := f(el)
		if val != nil {
			output = append(output, *val)
		}
	}
	return output
}

// Returns the first element for which the function
// returns a true value.
func Find[T any](input []T, f func(T) bool) *T {
	for _, el := range input {
		if f(el) {
			return &el
		}
	}
	return nil
}

func First[T any](input []T) *T {
	if len(input) > 0 {
		return &input[0]
	} else {
		return nil
	}
}

// Maps each input element of type T to type K using function f,
// where f returns []K
func FlatMap[T, K any](input []T, f func(T) []K) []K {
	output := make([]K, 0, len(input))
	for _, el := range input {
		output = append(output, f(el)...)
	}
	return output
}

// Returns a Map
// Each key is a return value from the function f
// Each value is a slice of those elements
// for which the function returned that key.
func GroupBy[T any, K comparable](input []T, f func(T) K) map[K][]T {
	output := make(map[K][]T)
	for _, el := range input {
		key := f(el)
		output[key] = append(output[key], el)
	}
	return output
}

func IsMember[T any](input []T, member T) bool {
	for _, el := range input {
		if reflect.DeepEqual(el, member) {
			return true
		}
	}
	return false
}

func Last[T any](input []T) *T {
	length := len(input)
	if length > 0 {
		return &input[length-1]
	} else {
		return nil
	}
}

// Maps each input element of type T to type K using function f
func Map[T, K any](input []T, f func(T) K) []K {
	output := make([]K, len(input))
	for i, el := range input {
		output[i] = f(el)
	}
	return output
}

// Return a new slice containing the first n elements
// of input
func Take[T any](input []T, n int) []T {
	return input[0:n]
}
