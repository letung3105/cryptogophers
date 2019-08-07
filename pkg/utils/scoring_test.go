package utils

import "testing"

func TestScoreTxtEn(t *testing.T) {
	t.Parallel()
	tt := []struct {
		input    []byte
		expected float64
	}{
		{input: []byte(""), expected: 0},
		{input: []byte("this is a test input"), expected: float64(93) / float64(20)},
		{input: []byte("gibberishdjfkljsdklf"), expected: float64(304) / float64(20)},
		{input: []byte("\xff\x47\xaf\xab\xff\x23\xed\xac\x04"), expected: float64(644) / float64(9)},
	}

	for _, tc := range tt {
		output := ScoreTxtEn(tc.input)
		if output != tc.expected {
			t.Errorf("Unexpected score: got %.6f, expected %.6g", output, tc.expected)
		}
	}
}
