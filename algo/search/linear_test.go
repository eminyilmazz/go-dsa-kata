package search

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		haystack []interface{}
		needle   interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty haystack",
			args: args{haystack: []interface{}{}, needle: 5},
			want: -1,
		},
		{
			name: "Element exists in the haystack",
			args: args{haystack: []interface{}{1, 2, 3, 4, 5}, needle: 3},
			want: 2,
		},
		{
			name: "Element does not exist in the haystack",
			args: args{haystack: []interface{}{1, 2, 3, 4, 5}, needle: 6},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
