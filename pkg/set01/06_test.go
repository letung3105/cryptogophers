package set01

import "testing"

func TestRepeatingXORDecrypt(t *testing.T) {
	t.Parallel()
	test := struct {
		filepath      string
		keysizeMax    int
		keysizeTrials int
	}{
		"../../data/06.txt", 40, 1,
	}

	out, key, score, err := RepeatingXORDecrypt(test.filepath, test.keysizeMax, test.keysizeTrials)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	t.Logf("key: %s | score: %.4f\n%s", key, score, out)
}
