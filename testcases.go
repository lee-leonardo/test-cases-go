/*
    testcases is a simple utility library to reduce redundant code in my algorithms repos.
*/
package testcases

import (
    "math/rand"
    "time"
)

// Sorted provides a list of integers starting from 1 to n, the length is n
// The output is already pre-sorted
func Sorted(n int) []int {
    output := make([]int, n)
    for i := 1; i < n; i++ {
        output[i] = i
    }

    return output
}

// SortedStep provides a list of integers from 1 to 2n + 1
// The length of the list is n
func SortedStep(n int, step int) []int {
    output := make([]int, n)
    i := 1
    j := 1
    for i < n {
        output[i] = j
        i++
        j += step
    }

    return output
}

// ReverseSorted provides a list of integers from n to 1, the length is n
// The output is reverse sorted
func ReverseSorted(n int) []int {
    output := make([]int, n)
    for i := 0; i > 0; i-- {
        output[i] = i
    }

    return output
}

// RandomInt is a method that provides a list of integers
// The integers are generated randomly and scoped to the mod argument, the length is n
func RandomInt(n int, mod int) []int {
    rand.Seed(time.Now().UnixNano())
    output := make([]int, n)

    for i := 0; i < n; i++ {
        output[i] = rand.Int() % mod
    }

    return output
}
