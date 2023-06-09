package match

import "testing"

func TestMatchExactly(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
		want     bool
	}{
		{
			name:     "nil slices",
			expected: nil,
			actual:   nil,
			want:     false,
		},
		{
			name:     "non slice types",
			expected: 1,
			actual:   "one",
			want:     false,
		},
		{
			name:     "different length slices",
			expected: []int{1, 2, 3},
			actual:   []int{1, 2, 3, 4},
			want:     false,
		},
		{
			name:     "different types in slices",
			expected: []int{1, 2, 3},
			actual:   []string{"1", "2", "3"},
			want:     false,
		},
		{
			name:     "different values in slices",
			expected: []int{1, 2, 3},
			actual:   []int{1, 2, 4},
			want:     false,
		},
		{
			name:     "same slices with different order",
			expected: []int{1, 2, 3},
			actual:   []int{3, 2, 1},
			want:     false,
		},
		{
			name:     "empty slices",
			expected: []int{},
			actual:   []int{},
			want:     true,
		},
		{
			name:     "same slices",
			expected: []int{1, 2, 3},
			actual:   []int{1, 2, 3},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchExactly(tt.expected, tt.actual); got != tt.want {
				t.Errorf("exactMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchElements(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
		want     bool
	}{
		{
			name:     "nil slices",
			expected: nil,
			actual:   nil,
			want:     false,
		},
		{
			name:     "non slice types",
			expected: 1,
			actual:   "one",
			want:     false,
		},
		{
			name:     "different length slices",
			expected: []int{1, 2, 3},
			actual:   []int{1, 2, 3, 4},
			want:     false,
		},
		{
			name:     "different types in slices",
			expected: []int{1, 2, 3},
			actual:   []string{"1", "2", "3"},
			want:     false,
		},
		{
			name:     "different values in slices",
			expected: []int{1, 2, 3},
			actual:   []int{1, 2, 4},
			want:     false,
		},
		{
			name:     "same slices with different order",
			expected: []int{1, 2, 3},
			actual:   []int{3, 2, 1},
			want:     true,
		},
		{
			name:     "same slices in the same order",
			expected: []int{1, 2, 3},
			actual:   []int{1, 2, 3},
			want:     true,
		},
		{
			name:     "empty slices",
			expected: []int{},
			actual:   []int{},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchElements(tt.expected, tt.actual); got != tt.want {
				t.Errorf("matchElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
