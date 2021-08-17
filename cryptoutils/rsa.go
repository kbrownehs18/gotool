package cryptoutils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RSAEncode RSA encode
func RSAEncode(data []byte, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// RSADecode RSA decode
func RSADecode(data []byte, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, data)
}

// RSABase64Encode RSA encode base64 to string
func RSABase64Encode(data []byte, publicKey []byte) (string, error) {
	b, err := RSAEncode(data, publicKey)
	if err != nil {
		return "", err
	}

	return Base64Encode(b), nil
}

// RSABase64Decode RSA decode to string
func RSABase64Decode(data string, publicKey []byte) (string, error) {
	b, err := Base64Decode(data)
	if err != nil {
		return "", err
	}
	bb, err := RSADecode(b, publicKey)
	if err != nil {
		return "", err
	}

	return string(bb), nil
}
