package p2_integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(5, 7)
	fmt.Println(sum)
	// Output: 12
}

func TestAdder(t *testing.T) {
	sum := Add(2, 4)
	expected := 6

	if sum != expected {
		t.Errorf("Got %d, want %d", sum, expected)
	}
}
