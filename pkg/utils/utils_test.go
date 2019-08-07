package utils

import (
	"reflect"
	"testing"
)

func TestBytesBlocksMake(t *testing.T) {
	t.Parallel()
	tt := []struct {
		input     []byte
		blocksize uint
		expected  [][]byte
	}{
		{
			input:     []byte("aaabbbccc"),
			blocksize: 3,
			expected:  [][]byte{[]byte("aaa"), []byte("bbb"), []byte("ccc")},
		},
		{
			input:     []byte("aa"),
			blocksize: 3,
			expected:  [][]byte{[]byte("aa")},
		},
		{
			input:     []byte("aaabbbcccd"),
			blocksize: 3,
			expected:  [][]byte{[]byte("aaa"), []byte("bbb"), []byte("ccc"), []byte("d")},
		},
		{
			input:     []byte(""),
			blocksize: 3,
		},
	}

	for _, tc := range tt {
		blocks := BytesBlockMake(tc.input, tc.blocksize)
		if !reflect.DeepEqual(blocks, tc.expected) {
			t.Errorf("Unexpected output: got %v, expected %v", blocks, tc.expected)
		}
	}
}

func TestBytesBlocksTranspose(t *testing.T) {
	t.Parallel()
	tt := []struct {
		input    [][]byte
		expected [][]byte
	}{
		{
			input:    [][]byte{[]byte("aaa"), []byte("bbb"), []byte("ccc")},
			expected: [][]byte{[]byte("abc"), []byte("abc"), []byte("abc")},
		},
		{
			input:    [][]byte{[]byte("aa")},
			expected: [][]byte{[]byte("a"), []byte("a")},
		},
		{
			input:    [][]byte{[]byte("aaa"), []byte("bbb"), []byte("ccc"), []byte("d")},
			expected: [][]byte{[]byte("abcd"), []byte("abc"), []byte("abc")},
		},
		{
			input:    [][]byte{[]byte("a")},
			expected: [][]byte{[]byte("a")},
		},
		{
			input: [][]byte{},
		},
	}

	for _, tc := range tt {
		blocks := BytesBlocksTranspose(tc.input)
		if !reflect.DeepEqual(blocks, tc.expected) {
			t.Errorf("Unexpected output: got %v, expected %v", blocks, tc.expected)
		}
	}
}
