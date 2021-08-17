package cryptoutils

import (
	"encoding/base64"
)

// Base64Encode string encode
func Base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Base64Decode string decode
func Base64Decode(str string) ([]byte, error) {
	x := len(str) * 3 % 4
	switch {
	case x == 2:
		str += "=="
	case x == 1:
		str += "="
	}
	return base64.StdEncoding.DecodeString(str)
}
