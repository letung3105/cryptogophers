package crypts

import (
	"encoding/hex"
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
			hasError: true,
		},
		{
			src:      []byte("\x01\x03\x05\x07\x09"),
			target:   []byte("\x00\x02\x04\x06"),
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
		} else {
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("Unexpected output: got %s, expected %s", output, tc.expected)
			}
		}
	}
}

func TestSingleByteXOR(t *testing.T) {
	t.Parallel()
	tt := []struct {
		src      []byte
		target   byte
		expected []byte
	}{
		{
			src:      []byte("\x01\x03\x05\x07\x09"),
			target:   byte(1),
			expected: []byte("\x00\x02\x04\x06\x08"),
		},
		{
			src:      []byte("anfktue"),
			target:   byte('z'),
			expected: []byte("\x1b\x14\x1c\x11\x0e\x0f\x1f"),
		},
		{
			src:      []byte(""),
			target:   byte(2),
			expected: []byte(""),
		},
	}
	for _, tc := range tt {
		output := SingleByteXOR(tc.src, tc.target)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("Unexpected output: got %s, expected %s", output, tc.expected)
		}
	}
}

func TestRepeatingXOR(t *testing.T) {
	t.Parallel()
	tt := []struct {
		plain     []byte
		key       []byte
		cipherHex []byte
	}{
		{
			plain:     []byte("This is a test plaintext"),
			key:       []byte("TEST"),
			cipherHex: []byte("002d3a27742c2074356527312731732438243a3a20202b20"),
		},
		{
			plain:     []byte("This is another test plaintext"),
			key:       []byte("T"),
			cipherHex: []byte("003c3d27743d2774353a3b203c31267420312720742438353d3a20312c20"),
		},
		{
			plain:     []byte("dkgtbkljrtljerkltjerkjfklsdjtiolnjbv"),
			key:       []byte("dkalcitk"),
			cipherHex: []byte("0000061801021801161f0d06061b1f071001041e080312000818050617001b070a01031a"),
		},
		{
			plain:     []byte("Short key!"),
			key:       []byte("A very long key"),
			cipherHex: []byte("1248191706594b09164f"),
		},
		{
			plain:     []byte("This test has no key"),
			key:       []byte(""),
			cipherHex: []byte("54686973207465737420686173206e6f206b6579"),
		},
	}

	for _, tc := range tt {
		output := RepeatingXOR(tc.plain, tc.key)
		cipher := make([]byte, hex.DecodedLen(len(tc.cipherHex)))
		n, err := hex.Decode(cipher, tc.cipherHex)
		if err != nil {
			t.Fatalf("Could not decode hex: %v", err)
		}
		if !reflect.DeepEqual(output, cipher[:n]) {
			t.Errorf("Unexpected output: got %s, expected %s", output, cipher[:n])
		}
	}
}
