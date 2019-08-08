package set01

import "testing"

func TestDectectSingleByteXOR(t *testing.T) {
	t.Parallel()
	filepath := "../../data/04.txt"

	plain, key, err := DetectSingleByteXOR(filepath)
	if err != nil {
		t.Fatalf("Could not detect single byte xor cipher: %v", err)
	}
	t.Logf("Key: %c | Detected single byte XOR plaintext: %s", key, plain)
}
