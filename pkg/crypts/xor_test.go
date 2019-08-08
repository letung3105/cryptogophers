package crypts

import (
	"bytes"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		hasError bool
		in       []byte
		key      []byte
		out      []byte
	}{
		{
			"EqualLength",
			false,
			[]byte{0x01, 0x03, 0x05, 0x07, 0x09},
			[]byte{0x00, 0x02, 0x04, 0x06, 0x08},
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
			[]byte{0x01, 0x03, 0x05, 0x07},
			[]byte{0x00, 0x02, 0x04, 0x06, 0x08},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := FixedXOR(test.in, test.key)
			if err != nil && !test.hasError {
				t.Fatalf("unexpected error: %+v", err)
			}
			if !bytes.Equal(out, test.out) {
				t.Errorf(
					"FixedXOR(%x, %x)\nhave %x\nwant %x",
					test.in, test.key, out, test.out,
				)
			}
		})
	}
}

func TestSingleByteXOR(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   []byte
		key  byte
		out  []byte
	}{
		{
			"InputFull",
			[]byte{0x01, 0x03, 0x05, 0x07, 0x09},
			byte(0x01),
			[]byte("\x00\x02\x04\x06\x08"),
		},
		{
			"InputEmpty",
			[]byte{},
			byte(0x00),
			[]byte{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := SingleByteXOR(test.in, test.key)
			if !bytes.Equal(out, test.out) {
				t.Errorf(
					"SingleByteXOR(%x, %x)\nhave%x\nwant %x",
					test.in, test.key, out, test.out,
				)
			}
		})
	}
}

func TestRepeatingXOR(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   []byte
		key  []byte
		out  []byte
	}{
		{
			"InputFullBlocks",
			[]byte{0x01, 0x02, 0x03, 0x01, 0x02, 0x03},
			[]byte{0x01, 0x02, 0x03},
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			"InputNotFullBlocks",
			[]byte{0x01, 0x02, 0x03, 0x04, 0x01, 0x02},
			[]byte{0x01, 0x02, 0x03, 0x04},
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			"InputSmallerThanKey",
			[]byte{0x01, 0x02, 0x03},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x01, 0x02},
			[]byte{0x00, 0x00, 0x00},
		},
		{
			"KeySingle",
			[]byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
			[]byte{0x01},
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			"InputEmpty",
			[]byte{},
			[]byte{0x01, 0x02, 0x03},
			[]byte{},
		},
		{
			"KeyEmpty",
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
			[]byte{},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := RepeatingXOR(test.in, test.key)
			if !bytes.Equal(out, test.out) {
				t.Errorf(
					"RepeatingXOR(%x, %x)\nhave%x\nwant %x",
					test.in, test.key, out, test.out,
				)
			}
		})
	}
}
