package utils

import (
	"reflect"
	"testing"
)

func TestBytesBlocksMake(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name      string
		blocksize int
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

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := BytesBlockMake(tc.in, tc.blocksize)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("unexpected output:\nhave %v\nwant %v", out, tc.out)
			}
		})
	}
}

func TestBytesBlocksTranspose(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name string
		in   [][]byte
		out  [][]byte
	}{
		{
			"EvenBlocks",
			[][]byte{[]byte("aaa"), []byte("bbb"), []byte("ccc")},
			[][]byte{[]byte("abc"), []byte("abc"), []byte("abc")},
		},
		{
			"OneBlock",
			[][]byte{[]byte("aa")},
			[][]byte{[]byte("a"), []byte("a")},
		},
		{
			"TrailingBlock",
			[][]byte{[]byte("aaa"), []byte("bbb"), []byte("ccc"), []byte("d")},
			[][]byte{[]byte("abcd"), []byte("abc"), []byte("abc")},
		},
		{
			"Singular",
			[][]byte{[]byte("a")},
			[][]byte{[]byte("a")},
		},
		{
			"Empty",
			[][]byte{},
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := BytesBlocksTranspose(tc.in)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("unexpected output:\nhave %v\nwant %v", out, tc.out)
			}
		})
	}
}

func TestHasDupBlock(t *testing.T) {
	t.Parallel()
	tt := []struct {
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

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := HasNonOverlapDup(tc.in, tc.blocksize)
			if out != tc.out {
				t.Errorf("unexpected output:\nhave %t\nwant %t", out, tc.out)
			}
		})
	}
}
