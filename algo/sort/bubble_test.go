package sort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Empty array",
			args: []int{},
			want: []int{},
		},
		{
			name: "Unsorted array",
			args: []int{5, 2, 8, 1, 0},
			want: []int{0, 1, 2, 5, 8},
		},
		{
			name: "Sorted array",
			args: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
