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
    for i := 0; i < n; i++ {
        output[i] = i + 1
    }

    return output
}

// SortedStep provides a list of integers from 1 to 2n + 1
// The length of the list is n
func SortedStep(n int, step int) []int {
    output := make([]int, n)
    for i, j := 0, 1; i < n; i, j = i + 1, j + step {
        output[i] = j
    }

    return output
}

// ReverseSorted provides a list of integers from n to 1, the length is n
// The output is reverse sorted
func ReverseSorted(n int) []int {
    output := make([]int, n)
    for i := n; i > 0; i-- {
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

// ScrambleList randomly swaps length of list times
func ScrambleList(list []int) {
    rand.Seed(time.Now().UnixNano())

    length := len(list)
    for i := range list {
        j := rand.Int() % length

        list[i], list[j] = list[j], list[i]
    }
}
