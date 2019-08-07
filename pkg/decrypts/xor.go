package decrypts

import (
	"math"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
)

// SingleByteXOR guesses plain text from cipher text by bruteforcing the key
func SingleByteXOR(cipher []byte) []byte {
	minScore := math.MaxFloat64
	var plain []byte
	for k := 0x00; k <= 0xff; k++ {
		dst := crypts.SingleByteXOR(cipher, byte(k))
		score := utils.ScoreTxtEn(dst)
		if score < minScore {
			minScore = score
			plain = dst
		}
	}
	return plain
}
