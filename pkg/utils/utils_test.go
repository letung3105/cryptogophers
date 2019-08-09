package utils

import (
	"reflect"
	"testing"
)

func TestBytesBlocksMake(t *testing.T) {
	t.Parallel()
	tests := []struct {
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := BytesBlockMake(test.in, test.blocksize)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("unexpected output:\nhave %v\nwant %v", out, test.out)
			}
		})
	}
}

func TestBytesBlocksTranspose(t *testing.T) {
	t.Parallel()
	tests := []struct {
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := BytesBlocksTranspose(test.in)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("unexpected output:\nhave %v\nwant %v", out, test.out)
			}
		})
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
				t.Errorf("unexpected output:\nhave %t\nwant %t", out, test.out)
			}
		})
	}
}

func TestIsEqualHex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		hasError  bool
		targetHex []byte
		in        []byte
		out       bool
	}{
		{
			"EqualHex",
			false,
			[]byte("5468697320697320612074657374"),
			[]byte("This is a test"),
			true,
		},
		{
			"NotEqualHex",
			false,
			[]byte("5468697320697320612074657374"),
			[]byte("This is a untrue test"),
			false,
		},
		{
			"InputEmpty",
			false,
			[]byte("5468697320697320612074657374"),
			[]byte(""),
			false,
		},
		{
			"HexEmpty",
			false,
			[]byte(""),
			[]byte("This is a test"),
			false,
		},
		{
			"BothEmpty",
			false,
			[]byte(""),
			[]byte(""),
			true,
		},
		{
			"InvalidHex",
			true,
			[]byte("5468697320697320612074657374a"),
			[]byte("This is a test"),
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := IsEqualHex(test.in, test.targetHex)
			if err != nil && !test.hasError {
				t.Fatalf("unexpected error: %+v", err)
			}
			if out != test.out {
				t.Errorf("unexpected output:\nhave %t\nwant %t", out, test.out)
			}
		})
	}
}

func TestIsEqualB64(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		hasError  bool
		targetB64 []byte
		in        []byte
		out       bool
	}{
		{
			"EqualB64",
			false,
			[]byte("VGhpcyBpcyBhIHRlc3Q="),
			[]byte("This is a test"),
			true,
		},
		{
			"NotEqualB64",
			false,
			[]byte("VGhpcyBpcyBhIHRlc3Q="),
			[]byte("This is a untrue test"),
			false,
		},
		{
			"InputEmpty",
			false,
			[]byte("VGhpcyBpcyBhIHRlc3Q="),
			[]byte(""),
			false,
		},
		{
			"B64Empty",
			false,
			[]byte(""),
			[]byte("This is a test"),
			false,
		},
		{
			"BothEmpty",
			false,
			[]byte(""),
			[]byte(""),
			true,
		},
		{
			"InvalidB64",
			true,
			[]byte("VGhpcyBpcyBhIHRlc3Q-="),
			[]byte("This is a test"),
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := IsEqualB64(test.in, test.targetB64)
			if err != nil && !test.hasError {
				t.Fatalf("unexpected error: %+v", err)
			}
			if out != test.out {
				t.Errorf("unexpected output:\nhave %t\nwant %t", out, test.out)
			}
		})
	}
}
