package utils

import "testing"

func TestHammingDistance(t *testing.T) {
	t.Parallel()
	tt := []struct {
		src      []byte
		target   []byte
		expected int
		hasError bool
	}{
		{
			src:      []byte("this is a test"),
			target:   []byte("wokka wokka!!!"),
			expected: 37,
			hasError: false,
		},
		{
			src:      []byte("this is a test"),
			target:   []byte("a unequal buffer"),
			expected: -1,
			hasError: true,
		},
		{
			src:      []byte("chewbacca"),
			target:   []byte("chewbacca"),
			expected: 0,
			hasError: false,
		},
		{
			src:      []byte(""),
			target:   []byte(""),
			expected: 0,
			hasError: false,
		},
	}

	for _, tc := range tt {
		output, err := HammingDistance(tc.src, tc.target)
		if err != nil && !tc.hasError {
			t.Errorf("Unexpected error: %v", err)
		} else {
			if output != tc.expected {
				t.Errorf("Unexpected output: got %d, expected %d", output, tc.expected)
			}
		}
	}
}
