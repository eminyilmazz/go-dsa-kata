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

func TestLinkedList_Prepend(t *testing.T) {
	l := &LinkedList{}
	v := 1
	l.Prepend(v)
	if l.head == nil {
		t.Errorf("Expected list head to be non-nil after prepend")
	}
	if l.head.val != v {
		t.Errorf("Expected list head value to be '%d', but got '%d'", v, l.head.val)
	}

	// Okay to be fair, unit tests can get really boring very fast. You're witnessing one.
	existingValue := l.head.val
	v = 420
	l.Prepend(v)
	if l.head == nil {
		t.Errorf("Expected list head to be non-nil after prepend")
	} else if l.head.val != v {
		t.Errorf("Expected list head value to be '%d', but got '%d'", v, l.head.val)
	} else if l.head.next.val != existingValue {
		t.Errorf("Expected list head's next node value to be '%d', but got '%d'", existingValue, l.head.next.val)
	}
}

func TestLinkedList_Append(t *testing.T) {
	l := &LinkedList{}

	veryCoolValue := 420
	l.Append(veryCoolValue)

	if l.head == nil {
		t.Errorf("Expected list head to be non-nil after prepend")
	}
	if l.head.val != veryCoolValue {
		t.Errorf("Expected list head value to be '%d', but got '%d'", veryCoolValue, l.head.val)
	}

	// you might have realized I'm just copy & pasting most of the stuff. This is why it's boring. But it looks good on my github
	// like any one reads this crap.

	anotherFancyValue := 1337
	l.Append(anotherFancyValue)
	current := l.head
	for current.next != nil {
		current = current.next
	}

	if current.val != anotherFancyValue {
		t.Errorf("Expected last node value to be '%d', but got '%d'", anotherFancyValue, current.val)
	} else if current.next != nil {
		t.Errorf("Expected last node's next pointer to be nil, but got non-nil")
	}
}
