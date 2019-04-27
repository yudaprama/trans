package closest

import "sort"

func SearchClosest(target float64, arr []float64) float64 {
	n := sort.SearchFloat64s(arr, target)
	if arr[n]-target <= target-arr[n-1] {
		n += 1
	}
	return arr[n-1]
}
