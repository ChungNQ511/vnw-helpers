package funcvx

// Ternary is a function that returns a value based on a condition
// It's useful for writing concise code
// @param cond bool - The condition to check
// @param a T - The value to return if the condition is true
// @param b T - The value to return if the condition is false
// @return T - The value to return
func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
