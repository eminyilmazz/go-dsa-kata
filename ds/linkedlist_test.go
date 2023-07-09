package ds

import (
	"errors"
	"testing"
)

func TestLinkedList_Get(t *testing.T) {
	l := &LinkedList{}
	l.Append(1)
	l.Append(420)
	l.Append(1337)

	// Test valid index
	expected := 420
	value, err := l.Get(1)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if value != expected {
		t.Errorf("Expected value '%d', but got: '%d'", expected, value)
	}

	// Test invalid index
	invalidIndex := uint(5)
	_, err = l.Get(invalidIndex)
	expectedError := errors.New("index out of bounds")
	if err == nil {
		t.Errorf("Expected error for index %d, but got no error", invalidIndex)
	} else if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%s', but got: '%s'", expectedError, err)
	}
}

func TestLinkedList_Length(t *testing.T) {
	l := &LinkedList{}
	expected := uint(0)
	actual := l.Length()

	if actual != expected {
		t.Errorf("Expected length %d, but got %d", expected, actual)
	}

	l.Append(1)
	l.Append(420)
	l.Append(1337)

	expected = uint(3)
	actual = l.Length()

	if actual != expected {
		t.Errorf("Expected length %d, but got %d", expected, actual)
	}
}

func TestLinkedList_Search(t *testing.T) {
	l := &LinkedList{}
	needle := "420"
	expected := -1
	actual := l.Search(needle)

	if expected != actual {
		t.Errorf("Expected index %d, but got %d", expected, actual)
	}

	l.Append("1")
	l.Append("420")
	l.Append("1337")
	expected = 1
	actual = l.Search(needle)

	if expected != actual {
		t.Errorf("Expected index %d, but got %d", expected, actual)
	}

	needle = "help me"
	expected = -1
	actual = l.Search(needle)

	if expected != actual {
		t.Errorf("Expected index %d, but got %d", expected, actual)
	}
}
