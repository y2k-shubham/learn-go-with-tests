package p4_arrays_n_slices

import "testing"

func TestSum(t *testing.T) {
	arr := []int{8, -1, 3, 7, -4, -2, 9}

	got := Sum(arr)
	want := 20

	if got != want {
		t.Errorf("got %d, want %d, given %v", got, want, arr)
	}
}
