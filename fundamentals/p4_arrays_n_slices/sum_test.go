package p4_arrays_n_slices

import "testing"

func assertHelper(t testing.TB, got, want int, slc []int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d, given %v", got, want, slc)
	}
}

func TestSum(t *testing.T) {
	t.Run("Slice (array of any len)", func(t *testing.T) {
		arr := []int{8, -1, 3, 7, -4, -2, 9}

		got := Sum(arr)
		want := 20

		assertHelper(t, got, want, arr)
	})
}
