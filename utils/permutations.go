// Package utils provides some useful tools for IA04
package utils

// ReverseIntSlice reverses the given slice
func ReverseIntSlice(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

// Factorial returns the product of all positive integers less than or equal to n
func Factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// FirstPermutation returns the slice of int [0, 1, ..., n-1]
func FirstPermutation(n int) (a []int) {
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	return
}

// NextPermutation returns the next lexicographical permutation of an integer slice
// The result is stored in a new slice
func NextPermutation(a []int) (res []int, ok bool) {
	n := len(a)
	res = make([]int, n)
	copy(res, a)

	if n < 2 {
		return res, false
	}

	k := n - 2
	for a[k] >= a[k+1] {
		k--
		if k < 0 {
			return res, false
		}
	}

	l := n - 1
	for a[k] >= a[l] {
		l--
	}

	res[k], res[l] = res[l], res[k]
	ReverseIntSlice(res[k+1:])

	return res, true
}
