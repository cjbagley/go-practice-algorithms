package twocrystalballsproblem

import (
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	const totalEmements int = 10000
	errorMsg := "Expected %d', got %d"
	haystack := make([]bool, totalEmements)
	index := rand.Intn(totalEmements)

	for i := index; i < totalEmements; i++ {
		haystack[i] = true
	}

	expected := index
	got := twoCrystalBalls(haystack)
	if expected != got {
		t.Errorf(errorMsg, expected, got)
	}
}
