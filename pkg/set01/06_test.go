package set01

import "testing"

func TestRepeatingXORDecrypt(t *testing.T) {
	t.Parallel()
	filepath := "../../data/06.txt"

	plain, key, err := RepeatingXORDecrypt(filepath, 40, 1)
	if err != nil {
		t.Fatalf("Could not detect single byte xor cipher: %v", err)
	}
	t.Logf("Key: %s | Repeating XOR plaintext: %s", key, plain)
}
