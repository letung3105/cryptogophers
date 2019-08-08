package utils

import "testing"

func TestHammingDistance(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		hasError bool
		target   []byte
		in       []byte
		out      int
	}{
		{
			"MatchLength",
			false,
			[]byte("wokka wokka!!!"),
			[]byte("this is a test"),
			37,
		},
		{
			"MismatchLength",
			true,
			[]byte("a unequal buffer"),
			[]byte("this is a test"),
			-1,
		},
		{
			"SameInputs",
			false,
			[]byte("chewbacca"),
			[]byte("chewbacca"),
			0,
		},
		{
			"EmptyInputs",
			false,
			[]byte(""),
			[]byte(""),
			0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := HammingDistance(test.in, test.target)
			if err != nil && !test.hasError {
				t.Fatalf("unexpected error: %+v", err)
			}
			if out != test.out {
				t.Errorf("unexpected output:\nhave %d\nwant %d", out, test.out)
			}
		})
	}
}
