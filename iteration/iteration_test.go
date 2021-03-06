package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	num := 5

	got := Repeat("a", num)
	expect := "aaaaa"
	if got != expect {
		t.Errorf("got %q but want %q", got, expect)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 9)
	}
}

func ExampleRepeat() {
	str := Repeat("qw", 3)
	fmt.Println(str)
	// Output: qwqwqw
}
