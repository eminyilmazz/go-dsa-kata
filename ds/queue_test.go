package ds

import (
	"testing"
)

func TestQueue_Dequeue(t *testing.T) {
	q := &Queue{}

	var expected interface{} = nil
	actual := q.Dequeue()

	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}

	q.elements = append(q.elements, 10, 3.14)
	expected = 10
	actual = q.Dequeue()

	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	} else if len(q.elements) != 1 {
		t.Errorf("Expected queue length '1' after but got '%d'", len(q.elements))
	}
	expected = 3.14
	actual = q.Dequeue()

	if actual != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	} else if len(q.elements) != 0 {
		t.Errorf("Expected queue length '0' after but got '%d'", len(q.elements))
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	type args struct {
		vars []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Enqueue existing queue",
			fields: fields{
				elements: []interface{}{1, 2, 3},
			},
			args: args{
				vars: []interface{}{"four", 5.0},
			},
		},
		{
			name: "Enqueue empty",
			fields: fields{
				elements: []interface{}{},
			},
			args: args{
				vars: []interface{}{1, 2, 3},
			},
		},
		{
			name: "Enqueue both empty",
			fields: fields{
				elements: []interface{}{},
			},
			args: args{
				vars: []interface{}{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				elements: tt.fields.elements,
			}
			q.Enqueue(tt.args.vars...)

			// Check the length of the queue after enqueueing
			expectedLength := len(tt.fields.elements) + len(tt.args.vars)
			actualLength := len(q.elements)
			if actualLength != expectedLength {
				t.Errorf("Expected queue length %d, but got %d", expectedLength, actualLength)
			}

			// Check the elements of the queue
			expectedElements := append(tt.fields.elements, tt.args.vars...)
			for i, expected := range expectedElements {
				actual := q.elements[i]
				if actual != expected {
					t.Errorf("Expected element at index %d to be %v, but got %v", i, expected, actual)
				}
			}
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Empty queue",
			fields: fields{
				elements: []interface{}{},
			},
			want: true,
		},
		{
			name: "Non-empty queue",
			fields: fields{
				elements: []interface{}{10, "Hello", 3.14},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				elements: tt.fields.elements,
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Size(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Queue size 0",
			fields: fields{
				elements: []interface{}{},
			},
			want: 0,
		},
		{
			name: "Queue size 3",
			fields: fields{
				elements: []interface{}{10, "Hello", 3.14},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				elements: tt.fields.elements,
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
