package twocrystalballsproblem

import (
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	// Test for finding correct index
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
		t.Errorf("Test 1: "+errorMsg, expected, got)
	}

	// Test to handle all results being false
	haystack = make([]bool, totalEmements)
	expected = -1
	got = twoCrystalBalls(haystack)
	if expected != got {
		t.Errorf("Test 2:"+errorMsg, expected, got)
	}

	// Test to handle empty haystack
	haystack = make([]bool, 0)
	expected = -1
	got = twoCrystalBalls(haystack)
	if expected != got {
		t.Errorf("Test 3:"+errorMsg, expected, got)
	}
}
