package lottery

import (
	"math/rand"
	"sort"
)

// drawUnique returns count unique numbers in the range [1, max].
// It uses the rand.Perm function to generate a random permutation of numbers and selects the first count numbers from it.
// The result is sorted before being returned.
func drawUnique(count, maxNum int) []int {
	perm := rand.Perm(maxNum)
	result := make([]int, count)
	for i := range result {
		result[i] = perm[i] + 1
	}
	sort.Ints(result)
	return result
}
