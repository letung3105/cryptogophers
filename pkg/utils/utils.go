package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"

	"github.com/pkg/errors"
)

// BytesBlockMake splits the buffer into chunks of with max length of blocksize
func BytesBlockMake(src []byte, blocksize int) [][]byte {
	var blocks [][]byte
	if blocksize > 0 {
		for len(src) > blocksize {
			blocks = append(blocks, src[:blocksize])
			src = src[blocksize:]
		}
		if len(src) > 0 {
			blocks = append(blocks, src)
		}
	}
	return blocks
}

// BytesBlocksTranspose stacks all blocks vertically and makes blocks outof each column
func BytesBlocksTranspose(src [][]byte) [][]byte {
	if len(src) == 0 {
		return nil
	}

	blocks := make([][]byte, len(src[0]))
	for j := 0; j < len(src[0]); j++ {
		var block []byte
		for i := 0; i < len(src); i++ {
			if len(src[i]) > j {
				block = append(block, src[i][j])
			}
		}
		blocks[j] = block
	}
	return blocks
}

// HasNonOverlapDup checks if the buffer has duplicated non-overlap blocks of length blocksize
func HasNonOverlapDup(src []byte, blocksize int) bool {
	if blocksize > 0 {
		for len(src) > blocksize {
			if bytes.Contains(src[blocksize:], src[:blocksize]) {
				return true
			}
			src = src[blocksize:]
		}
	}
	return false
}

// IsEqualHex compare a buffer with a hex encoded buffer
func IsEqualHex(src, targetHex []byte) (bool, error) {
	target := make([]byte, len(targetHex))
	n, err := hex.Decode(target, targetHex)
	if err != nil {
		return false, errors.Wrapf(err, "could not decode: %s", targetHex)
	}
	target = target[:n]

	return bytes.Equal(src, target), nil
}

// IsEqualB64 compare a buffer with a base64 encoded buffer
func IsEqualB64(src, targetB64 []byte) (bool, error) {
	b64 := base64.StdEncoding

	target := make([]byte, len(targetB64))
	n, err := b64.Decode(target, targetB64)
	if err != nil {
		return false, errors.Wrapf(err, "could not decode: %s", targetB64)
	}
	target = target[:n]

	return bytes.Equal(src, target), nil
}
