package set02

import (
	"github.com/letung3105/cryptogophers/pkg/utils"
)

// PKCS7PadChallenge using implementation of PKCS7 padding from utils package
func PKCS7PadChallenge(src []byte, paddedLen int) []byte {
	return utils.PKCS7Pad(src, paddedLen)
}
