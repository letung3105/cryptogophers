package utils

import (
	"reflect"
	"testing"
)

func TestBytesBlocksMake(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		blocksize uint
		in        []byte
		out       [][]byte
	}{
		{
			"EvenBlock",
			3,
			[]byte("aaabbbccc"),
			[][]byte{[]byte("aaa"), []byte("bbb"), []byte("ccc")},
		},
		{
			"FirstBlockTrailing",
			3,
			[]byte("aa"),
			[][]byte{[]byte("aa")},
		},
		{
			"LastBlockTrailing",
			3,
			[]byte("aaabbbcccd"),
			[][]byte{[]byte("aaa"), []byte("bbb"), []byte("ccc"), []byte("d")},
		},
		{
			"EmptyInput",
			3,
			[]byte(""),
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := BytesBlockMake(test.in, test.blocksize)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("BytesBlockMake(%q, %d)\nhave %v\nwant %v",
					test.in, test.blocksize, out, test.out)
			}
		})
	}
}

func TestBytesBlocksTranspose(t *testing.T) {
	t.Parallel()
	tt := []struct {
		input    [][]byte
		expected [][]byte
	}{
		{
			[][]byte{[]byte("aaa"), []byte("bbb"), []byte("ccc")},
			[][]byte{[]byte("abc"), []byte("abc"), []byte("abc")},
		},
		{
			[][]byte{[]byte("aa")},
			[][]byte{[]byte("a"), []byte("a")},
		},
		{
			[][]byte{[]byte("aaa"), []byte("bbb"), []byte("ccc"), []byte("d")},
			[][]byte{[]byte("abcd"), []byte("abc"), []byte("abc")},
		},
		{
			[][]byte{[]byte("a")},
			[][]byte{[]byte("a")},
		},
		{
			[][]byte{},
			nil,
		},
	}

	for _, tc := range tt {
		blocks := BytesBlocksTranspose(tc.input)
		if !reflect.DeepEqual(blocks, tc.expected) {
			t.Errorf("Unexpected output: got %v, expected %v", blocks, tc.expected)
		}
	}
}

func TestHasDupBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		blocksize int
		in        []byte
		out       bool
	}{
		{
			"DupFirstBlock",
			3,
			[]byte("aaabbbaaa"),
			true,
		},
		{
			"DupSecondBlock",
			3,
			[]byte("aaabbbcccbbb"),
			true,
		},
		{
			"LargeBlockSize",
			5,
			[]byte("aaaaaa"),
			false,
		},
		{
			"ZeroBlockSize",
			0,
			[]byte("aaaaaaaa"),
			false,
		},
		{
			"EmptyInput",
			3,
			[]byte(""),
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := HasNonOverlapDup(test.in, test.blocksize)
			if out != test.out {
				t.Errorf(
					"Input: '%s' | Blocksize: %d\nhave %t\nwant %t",
					test.in, test.blocksize, out, test.out,
				)
			}
		})
	}
}
