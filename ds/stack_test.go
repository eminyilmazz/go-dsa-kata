package ds

import (
	"reflect"
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Empty stack",
			fields: fields{
				elements: []interface{}{},
			},
			want: true,
		},
		{
			name: "Non-empty stack",
			fields: fields{
				elements: []interface{}{10, "Hello", 3.14},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				elements: tt.fields.elements,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "Empty queue",
			fields: fields{
				elements: []interface{}{},
			},
			want: nil,
		},
		{
			name: "Non-empty queue",
			fields: fields{
				elements: []interface{}{10, "Hello", 3.14},
			},
			want: 3.14,
		},
		{
			name: "Single queue",
			fields: fields{
				elements: []interface{}{10},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				elements: tt.fields.elements,
			}
			if got := s.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "Empty stack",
			fields: fields{
				elements: []interface{}{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Non-empty stack",
			fields: fields{
				elements: []interface{}{420, "ayo", 3.14},
			},
			want:    3.14,
			wantErr: false,
		},
		{
			name: "Single stack",
			fields: fields{
				elements: []interface{}{"ayo"},
			},
			want:    "ayo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				elements: tt.fields.elements,
			}
			got, err := s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	type args struct {
		e interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantLength int
		wantTop    interface{}
	}{
		{
			name: "Empty stack",
			fields: fields{
				elements: []interface{}{},
			},
			args: args{
				e: 124124,
			},
			wantLength: 1,
			wantTop:    124124,
		},
		{
			name: "Non-empty stack",
			fields: fields{
				elements: []interface{}{42, 2},
			},
			args: args{
				e: "I LOVE GOLANG",
			},
			wantLength: 3,
			wantTop:    "I LOVE GOLANG",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				elements: tt.fields.elements,
			}
			s.Push(tt.args.e)
			if s.Size() != tt.wantLength {
				t.Errorf("Expected stack length %d, but got %d", tt.wantLength, s.Size())
			}
			if top := s.Peek(); !reflect.DeepEqual(top, tt.wantTop) {
				t.Errorf("Expected top element %v, but got %v", tt.wantTop, top)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	type fields struct {
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Empty stack",
			fields: fields{
				elements: []interface{}{},
			},
			want: 0,
		},
		{
			name: "Non-empty stack",
			fields: fields{
				elements: []interface{}{420, "ayo", 3.14},
			},
			want: 3,
		},
		{
			name: "Single stack",
			fields: fields{
				elements: []interface{}{"ayo"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				elements: tt.fields.elements,
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
