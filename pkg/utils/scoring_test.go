package utils

import "testing"

func TestScoreTxtEn(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   []byte
		out  float64
	}{
		{
			"Empty",
			[]byte(""),
			0,
		},
		{
			"English",
			[]byte("this is a test input"),
			float64(93) / float64(20),
		},
		{
			"Nonsense",
			[]byte("gibberishdjfkljsdklf"),
			float64(304) / float64(20),
		},
		{
			"HexString",
			[]byte("\xff\x47\xaf\xab\xff\x23\xed\xac\x04"),
			float64(644) / float64(9),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := ScoreTxtEn(test.in)
			if out != test.out {
				t.Errorf("unexpected output:\nhave %.4f\nwant %.4f", out, test.out)
			}
		})
	}
}
