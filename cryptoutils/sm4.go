package cryptoutils

import (
	log "github.com/sirupsen/logrus"
	"github.com/tjfoc/gmsm/sm4"
)

func Sm4Encode(data string, key []byte) (string, error) {
	b, err := sm4.Sm4Ecb(key, []byte(data), true)
	if err != nil {
		log.Error("Sm4Encode error: ", err)
		return "", err
	}

	return Base64Encode(b), nil
}

func Sm4Decode(data string, key []byte) (string, error) {
	b, err := Base64Decode(data)
	if err != nil {
		log.Error("Sm4Decode base64 error: ", err)
		return "", err
	}
	b, err = sm4.Sm4Ecb(key, b, false)
	if err != nil {
		log.Error("Sm4Decode error: ", err)
		return "", err
	}

	return string(b), nil
}
