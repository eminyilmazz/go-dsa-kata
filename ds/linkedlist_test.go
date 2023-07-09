package ds

import (
	"errors"
	"testing"
)

func TestLinkedList_Get(t *testing.T) {
	list := &LinkedList{}
	list.Append(1)
	list.Append(420)
	list.Append(1337)

	// Test valid index
	expected := 420
	value, err := list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if value != expected {
		t.Errorf("Expected value '%d', but got: '%d'", expected, value)
	}

	// Test invalid index
	invalidIndex := uint(5)
	_, err = list.Get(invalidIndex)
	expectedError := errors.New("index out of bounds")
	if err == nil {
		t.Errorf("Expected error for index %d, but got no error", invalidIndex)
	} else if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%s', but got: '%s'", expectedError, err)
	}
}
