package cowabunga

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

// Maps each input element of type T to type K using function f
func Map[T, K any](input []T, f func(T) K) []K {
	output := make([]K, len(input))
	for i, el := range input {
		output[i] = f(el)
	}
	return output
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
