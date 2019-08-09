package crypts

import (
	"bytes"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name     string
		hasError bool
		key      []byte
		in       []byte
		out      []byte
	}{
		{
			"EqualLength",
			false,
			[]byte{0x00, 0x02, 0x04, 0x06, 0x08},
			[]byte{0x01, 0x03, 0x05, 0x07, 0x09},
			[]byte{0x01, 0x01, 0x01, 0x01, 0x01},
		},
		{
			"EmptyInputs",
			false,
			[]byte{},
			[]byte{},
			[]byte{},
		},
		{
			"MismatchLength",
			true,
			[]byte{0x00, 0x02, 0x04, 0x06, 0x08},
			[]byte{0x01, 0x03, 0x05, 0x07},
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out, err := FixedXOR(tc.in, tc.key)
			if err != nil && !tc.hasError {
				t.Fatalf("unexpected error: %+v", err)
			}
			if !bytes.Equal(out, tc.out) {
				t.Errorf("unexpected output:\nhave %x\nwant %x", out, tc.out)
			}
		})
	}
}

func TestSingleXOR(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name string
		key  byte
		in   []byte
		out  []byte
	}{
		{
			"InputFull",
			byte(0x01),
			[]byte{0x01, 0x03, 0x05, 0x07, 0x09},
			[]byte("\x00\x02\x04\x06\x08"),
		},
		{
			"InputEmpty",
			byte(0x00),
			[]byte{},
			[]byte{},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := SingleXOR(tc.in, tc.key)
			if !bytes.Equal(out, tc.out) {
				t.Errorf("unexpected output:\nhave %x\nwant %x", out, tc.out)
			}
		})
	}
}

func TestRepeatingXOR(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name string
		key  []byte
		in   []byte
		out  []byte
	}{
		{
			"InputFullBlocks",
			[]byte{0x01, 0x02, 0x03},
			[]byte{0x01, 0x02, 0x03, 0x01, 0x02, 0x03},
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			"InputNotFullBlocks",
			[]byte{0x01, 0x02, 0x03, 0x04},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x01, 0x02},
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			"InputSmallerThanKey",
			[]byte{0x01, 0x02, 0x03, 0x04, 0x01, 0x02},
			[]byte{0x01, 0x02, 0x03},
			[]byte{0x00, 0x00, 0x00},
		},
		{
			"KeySingle",
			[]byte{0x01},
			[]byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			"InputEmpty",
			[]byte{0x01, 0x02, 0x03},
			[]byte{},
			[]byte{},
		},
		{
			"KeyEmpty",
			[]byte{},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := RepeatingXOR(tc.in, tc.key)
			if !bytes.Equal(out, tc.out) {
				t.Errorf("unexpected output:\nhave %x\nwant %x", out, tc.out)
			}
		})
	}
}
