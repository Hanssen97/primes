package primes

// Merge from https://austingwalters.com/merge-sort-in-go-golang/ (modified)
func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	lenL, lenR := len(left)-1, len(right)-1

	for k := 0; k < size; k++ {
		if i > lenL && j <= lenR {
			slice[k] = right[j]
			j++
		} else if j > lenR && i <= lenL {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}