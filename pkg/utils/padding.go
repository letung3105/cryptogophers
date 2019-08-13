package utils

import "bytes"

// PKCS7Pad pads the source buffer to specified length according the the PKCS#7 specifications
func PKCS7Pad(src []byte, paddedLen int) []byte {
	padChar := paddedLen - len(src)
	if padChar <= 0 || padChar > 255 {
		return src
	}

	dst := make([]byte, paddedLen)
	copy(dst, src)
	pad := bytes.Repeat([]byte{byte(padChar)}, padChar)
	copy(dst[len(src):], pad)

	return dst
}

// PKCS7Valid checks if the bytes slice has a valid PKCS#7 padding
func PKCS7Valid(src []byte) bool {
	if len(src) <= 0 {
		return false
	}
	padChar := src[len(src)-1]
	pad := bytes.Repeat([]byte{padChar}, int(padChar))
	return bytes.HasSuffix(src, pad)
}

// PKCS7Unpad removes padding characters from buffers
func PKCS7Unpad(src []byte) []byte {
	if !PKCS7Valid(src) {
		return src
	}

	padLen := int(src[len(src)-1])
	dst := make([]byte, len(src)-padLen)
	copy(dst, src[:len(src)-padLen])

	return dst
}
