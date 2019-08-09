package utils

// PKCS7Pad pads the source buffer to specified length according the the PKCS#7 specifications
func PKCS7Pad(src []byte, paddedLen int) []byte {
	lenDiff := paddedLen - len(src)
	if lenDiff <= 0 || lenDiff > 255 {
		return src
	}

	dst := make([]byte, paddedLen)
	copy(dst, src)
	padChar := byte(lenDiff)
	for i := len(src); i < paddedLen; i++ {
		dst[i] = padChar
	}
	return dst
}
