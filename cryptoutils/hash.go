package cryptoutils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// Md5Sum md5
func Md5Sum(text string) string {
	h := md5.New()
	io.WriteString(h, text)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 return string
func SHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// HMacSHA256 hmac sha256
func HMacSHA256(s, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA1 return string
func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
