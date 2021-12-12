package p4_arrays_n_slices

import (
	"reflect"
	"testing"
)

func assertHelper1(t testing.TB, got, want int, slc []int) {
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

		assertHelper1(t, got, want, arr)
	})
}

func assertHelper2(t testing.TB, got, want []int, slc [][]int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v, given %v", got, want, slc)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("Varying no of slices: 0 slices", func(t *testing.T) {
		slc := [][]int{}

		got := SumAll(slc)
		want := []int{}

		assertHelper2(t, got, want, slc)
	})

	t.Run("Varying no of slices: 1 slice", func(t *testing.T) {
		slc := [][]int{[]int{1, 2}}

		got := SumAll(slc)
		want := []int{3}

		assertHelper2(t, got, want, slc)
	})

	t.Run("Varying no of slices: 2 slices", func(t *testing.T) {
		slc := [][]int{[]int{0, 9, -5}, []int{3}}

		got := SumAll(slc)
		want := []int{4, 3}

		assertHelper2(t, got, want, slc)
	})

	t.Run("Varying no of slices: 3 slices", func(t *testing.T) {
		slc := [][]int{[]int{1, 2}, []int{0, 9, -5}, []int{3}}

		got := SumAll(slc)
		want := []int{3, 4, 3}

		assertHelper2(t, got, want, slc)
	})
}
