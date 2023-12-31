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

func TestLinkedList_InsertAt(t *testing.T) {
	// Create an empty linked list
	list := &LinkedList{}

	value := 420
	err := list.InsertAt(0, value)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if list.head == nil {
		t.Errorf("Expected list head to be non-nil after insert")
	} else if list.head.val != value {
		t.Errorf("Expected list head value to be '%d', but got '%d'", value, list.head.val)
	}

	existingValue := list.head.val
	value = 1337
	err = list.InsertAt(0, value)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if list.head == nil {
		t.Errorf("Expected list head to be non-nil after insert")
	} else if list.head.val != value {
		t.Errorf("Expected list head value to be '%d', but got '%d'", value, list.head.val)
	} else if list.head.next == nil {
		t.Errorf("Expected list head's next node to be non-nil after insert")
	} else if list.head.next.val != existingValue {
		t.Errorf("Expected list head's next node value to be '%d', but got '%d'", existingValue, list.head.next.val)
	}

	invalidIndex := uint(5)
	err = list.InsertAt(invalidIndex, 6969)
	expectedError := errors.New("index out of bounds")
	if err == nil {
		t.Errorf("Expected error for index %d, but got no error", invalidIndex)
	} else if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%s', but got: '%s'", expectedError, err)
	}
}

func TestLinkedList_RemoveTail(t *testing.T) {
	l := &LinkedList{}

	err := l.removeTail()
	expectedError := errors.New("empty linked list")
	if err == nil {
		t.Errorf("Expected error '%s', but got no error", expectedError)
	} else if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%s', but got: '%s'", expectedError, err)
	}

	l.Append(1)
	l.Append(420)
	l.Append(1337)

	err = l.removeTail()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	if current.val != 420 {
		t.Errorf("Expected new tail value to be '420', but got '%d'", current.val)
	}

	err = l.removeTail()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	current = l.head
	for current.next != nil {
		current = current.next
	}
	if current.val != 1 {
		t.Errorf("Expected new tail value to be '1', but got '%d'", current.val)
	}

	err = l.removeTail()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	// Verify the list becomes empty
	if l.head != nil {
		t.Errorf("Expected list head to be nil after removing tail, but got non-nil")
	}
}

func TestLinkedList_RemoveHead(t *testing.T) {
	l := &LinkedList{}

	err := l.removeHead()
	expectedError := errors.New("empty linked list")
	if err == nil {
		t.Errorf("Expected error '%s', but got no error", expectedError)
	} else if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%s', but got: '%s'", expectedError, err)
	}

	l.Append(1)
	l.Append(2)
	l.Append(3)

	err = l.removeHead()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if l.head == nil {
		t.Errorf("Expected list head to be non-nil after removeHead")
	}

	if l.head.val != 2 {
		t.Errorf("Expected new head value to be '2', but got '%d'", l.head.val)
	}

	err = l.removeHead()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if l.head == nil {
		t.Errorf("Expected list head to be non-nil after removeHead")
	}

	if l.head.val != 3 {
		t.Errorf("Expected new tail value to be '3', but got '%d'", l.head.val)
	}

	err = l.removeHead()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if l.head != nil {
		t.Errorf("Expected list head to be nil after removing head, but got non-nil")
	}
}

func TestLinkedList_RemoveAt(t *testing.T) {
	l := &LinkedList{}

	err := l.removeAt(0)
	expectedErr := errors.New("index out of bounds")

	if err == nil {
		t.Errorf("Expected error '%s' but got no error", expectedErr)
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error '%s', but got: '%s'", expectedErr, err)
	}

	l.Prepend("up")
	l.Prepend("whats")
	l.Prepend("ayo")

	err = l.removeAt(0)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if l.head.val != "whats" {
		t.Errorf("Expected new head value to be 'whats', but got '%s'", l.head.val)
	}

	l.Prepend("ayo")
	err = l.removeAt(1)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if l.Length() != 2 {
		t.Errorf("Expected list length to be 2 after removal, but got %d", l.Length())
	}
	if l.head.val != "ayo" {
		t.Errorf("Expected head value to be 'ayo', but got '%s'", l.head.val)
	}

	err = l.removeAt(2) //out of bounds

	if err == nil {
		t.Errorf("Expected no error, but got: %v", err)
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error '%s', but got: '%s'", expectedErr, err)
	}
}
