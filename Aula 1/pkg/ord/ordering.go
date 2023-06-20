package ord

import "sort"

func Ord(nums []int) []int {
	sort.Ints(nums[:])
	return nums
}
