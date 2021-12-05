package p3_iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("bmg")
	fmt.Println(repeated)
	// Output: bmgbmgbmgbmgbmg
}

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got %s but expected %s", got, want)
	}
}
