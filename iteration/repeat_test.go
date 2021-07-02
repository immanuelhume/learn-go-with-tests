package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("yo", 4)
	fmt.Println(repeated)
	// Output: yoyoyoyo
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestContains(t *testing.T) {
	result := strings.Contains("doge", "ge")
	expected := true
	if expected != result {
		t.Errorf("expected %t but got %t", expected, result)
	}
}
