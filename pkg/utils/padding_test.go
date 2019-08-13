package utils

import (
	"bytes"
	"testing"
)

func TestPKCS7Pad(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name      string
		paddedLen int
		in        []byte
		out       []byte
	}{
		{
			"PadTo8",
			8,
			[]byte{0x01, 0x02, 0x03, 0x04},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x04, 0x04, 0x04, 0x04},
		},
		{
			"PadTo16",
			16,
			[]byte{0x01, 0x02, 0x03, 0x04},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c},
		},
		{
			"PadTo32",
			32,
			[]byte{0x01, 0x02, 0x03, 0x04},
			[]byte{
				0x01, 0x02, 0x03, 0x04, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c,
				0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c,
			},
		},
		{
			"InputEqualPaddedLen",
			8,
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		},
		{
			"InputLargerPaddedLen",
			4,
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		},
		{
			"PaddedLenExceedLimit",
			300,
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		},
		{
			"EmptyInput",
			8,
			[]byte{},
			[]byte{0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08},
		},
		{
			"PaddedLenZero",
			0,
			[]byte{0x00, 0x01, 0x02, 0x04},
			[]byte{0x00, 0x01, 0x02, 0x04},
		},
		{
			"NULL",
			0,
			[]byte{},
			[]byte{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := PKCS7Pad(tc.in, tc.paddedLen)
			if !bytes.Equal(out, tc.out) {
				t.Errorf("unexpected output:\nhave %x\nwant %x", out, tc.out)
			}
		})
	}
}

func TestPKCS7Valid(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name string
		in   []byte
		out  bool
	}{
		{
			"Pad4",
			[]byte{0x01, 0x02, 0x03, 0x04, 0x04, 0x04, 0x04, 0x04},
			true,
		},
		{
			"Pad12",
			[]byte{0x01, 0x02, 0x03, 0x04, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c},
			true,
		},
		{
			"Pad28",
			[]byte{
				0x01, 0x02, 0x03, 0x04, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c,
				0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c,
			},
			true,
		},
		{
			"InvalidPadding",
			[]byte{0x01, 0x02, 0x03, 0x04, 0x04, 0x03, 0x04, 0x04},
			false,
		},
		{
			"NoPadding",
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			false,
		},
		{
			"EmptyInput",
			[]byte{},
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := PKCS7Valid(tc.in)
			if out != tc.out {
				t.Errorf("unexpected output:\nhave %t\nwant %t", out, tc.out)
			}
		})
	}
}

func TestPKCS7Unpad(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name string
		in   []byte
		out  []byte
	}{
		{
			"Unpad4",
			[]byte{0x01, 0x02, 0x03, 0x04, 0x04, 0x04, 0x04, 0x04},
			[]byte{0x01, 0x02, 0x03, 0x04},
		},
		{
			"Unpad12",
			[]byte{0x01, 0x02, 0x03, 0x04, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c, 0x0c},
			[]byte{0x01, 0x02, 0x03, 0x04},
		},
		{
			"Unpad28",
			[]byte{
				0x01, 0x02, 0x03, 0x04, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c,
				0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c, 0x1c,
			},
			[]byte{0x01, 0x02, 0x03, 0x04},
		},
		{
			"InvalidPadding",
			[]byte{0x01, 0x02, 0x03, 0x04, 0x04, 0x03, 0x04, 0x04},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x04, 0x03, 0x04, 0x04},
		},
		{
			"NoPadding",
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		},
		{
			"EmptyInput",
			[]byte{},
			[]byte{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := PKCS7Unpad(tc.in)
			if !bytes.Equal(out, tc.out) {
				t.Errorf("unexpected output:\nhave %x\nwant %x", out, tc.out)
			}
		})
	}
}
