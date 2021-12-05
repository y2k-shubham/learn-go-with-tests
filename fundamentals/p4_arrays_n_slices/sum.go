package p4_arrays_n_slices

func Sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}
