package ciphers

import (
	"reflect"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	t.Parallel()
	tt := []struct {
		src      []byte
		target   []byte
		expected []byte
		hasError bool
	}{
		{
			src:      []byte("\x01\x03\x05\x07\x09"),
			target:   []byte("\x00\x02\x04\x06\x08"),
			expected: []byte("\x01\x01\x01\x01\x01"),
			hasError: false,
		},
		{
			src:      []byte("anfktue"),
			target:   []byte("qorjvba"),
			expected: []byte("\x10\x01\x14\x01\x02\x17\x04"),
			hasError: false,
		},
		{
			src:      []byte("\x01\x03\x05\x07"),
			target:   []byte("\x00\x02\x04\x06\x08"),
			expected: nil,
			hasError: true,
		},
		{
			src:      []byte("\x01\x03\x05\x07\x09"),
			target:   []byte("\x00\x02\x04\x06"),
			expected: nil,
			hasError: true,
		},
		{
			src:      []byte(""),
			target:   []byte(""),
			expected: []byte(""),
			hasError: false,
		},
	}

	for _, tc := range tt {
		output, err := FixedXOR(tc.src, tc.target)
		if err != nil && !tc.hasError {
			t.Errorf("Unexpected error occurs: %v", err)
		}
		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("Unexpected output: got %s, expected %s", output, tc.expected)
		}
	}
}
