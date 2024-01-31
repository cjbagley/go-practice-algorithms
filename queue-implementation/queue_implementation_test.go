package queueimplementation

import (
	"testing"
)

func TestQueue(t *testing.T) {
	errorMsg := "Expected %d', got %d"

	queue := NewQueue[int]()
	queue.Enqueue(5)
	queue.Enqueue(7)
	queue.Enqueue(9)

	expected := 5
	got, ok := queue.Dequeue()
	if !ok {
		t.Errorf("Test 1: did not dequeue successfully")
	} else if got != expected {
		t.Errorf("Test 1: "+errorMsg, expected, got)
	}

	currentLength := queue.Length
	expectedLength := 2
	if expectedLength != currentLength {
		t.Errorf("Test 2: "+errorMsg, expectedLength, currentLength)
	}

	queue.Enqueue(11)

	expected = 7
	got, ok = queue.Dequeue()
	if !ok {
		t.Errorf("Test 3: did not dequeue successfully")
	} else if got != expected {
		t.Errorf("Test 3: "+errorMsg, expected, got)
	}

	expected = 9
	got, ok = queue.Peek()
	if !ok {
		t.Errorf("Test 4: did not peek successfully")
	} else if got != expected {
		t.Errorf("Test 4: "+errorMsg, expected, got)
	}

	expected = 9
	got, ok = queue.Dequeue()
	if !ok {
		t.Errorf("Test 5: did not dequeue successfully")
	} else if got != expected {
		t.Errorf("Test 5: "+errorMsg, expected, got)
	}

	currentLength = queue.Length
	expectedLength = 1
	if expectedLength != currentLength {
		t.Errorf("Test 6: "+errorMsg, expectedLength, currentLength)
	}

	expected = 11
	got, ok = queue.Dequeue()
	if !ok {
		t.Errorf("Test 7: did not dequeue successfully")
	} else if got != expected {
		t.Errorf("Test 7: "+errorMsg, expected, got)
	}

	expected = 0
	got, ok = queue.Dequeue()
	if ok {
		t.Errorf("Test 8: dequeue should not be okay")
	}
	if got != expected {
		t.Errorf("Test 8: "+errorMsg, expected, got)
	}

	expected = 0
	got, ok = queue.Peek()
	if ok {
		t.Errorf("Test 9: peek should not be okay")
	}
	if got != expected {
		t.Errorf("Test 9: "+errorMsg, expected, got)
	}

	queue.Enqueue(100)
	queue.Enqueue(101)
	queue.Enqueue(102)

	currentLength = queue.Length
	expectedLength = 3
	if expectedLength != currentLength {
		t.Errorf("Test 10: "+errorMsg, expectedLength, currentLength)
	}
}
