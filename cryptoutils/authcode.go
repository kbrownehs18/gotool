package cryptoutils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kbrownehs18/gotool/arrayutils"
)

// AuthCodeType auth code type
type AuthCodeType int

const (
	// ENCODE encode str
	ENCODE AuthCodeType = iota
	// DECODE decode str
	DECODE
)

// Authcode Discuz Authcode golang version
// params[0] encrypt/decrypt bool true：encrypt false：decrypt, default: false
// params[1] key
// params[2] expires time(second)
// params[3] dynamic key length
func Authcode(text string, params ...interface{}) (str string, err error) {
	l := len(params)

	isEncode := DECODE
	key := "abcdefghijklmnopqrstuvwxyz0123456789"
	expiry := 0
	cKeyLen := 8

	if l > 0 {
		isEncode = params[0].(AuthCodeType)
	}

	if l > 1 {
		key = params[1].(string)
	}

	if l > 2 {
		expiry = params[2].(int)
		if expiry < 0 {
			expiry = 0
		}
	}

	if l > 3 {
		cKeyLen = params[3].(int)
		if cKeyLen < 0 {
			cKeyLen = 0
		}
	}
	if cKeyLen > 32 {
		cKeyLen = 32
	}

	timestamp := time.Now().Unix()

	// md5sum key
	mKey := Md5Sum(key)

	// keyA encrypt
	keyA := Md5Sum(mKey[0:16])
	// keyB validate
	keyB := Md5Sum(mKey[16:])
	// keyC dynamic key
	var keyC string
	if cKeyLen > 0 {
		if isEncode == ENCODE {
			// encrypt generate a key
			keyC = Md5Sum(fmt.Sprint(timestamp))[32-cKeyLen:]
		} else {
			// decrypt get key from header of string
			keyC = text[0:cKeyLen]
		}
	}

	// generate encrypt/decrypt key
	cryptKey := keyA + Md5Sum(keyA+keyC)
	// key length
	keyLen := len(cryptKey)
	if isEncode == ENCODE {
		// The first 10 strings is expires time
		// 10-26 strings is validator strings
		var d int64
		if expiry > 0 {
			d = timestamp + int64(expiry)
		}
		text = fmt.Sprintf("%010d%s%s", d, Md5Sum(text + keyB)[0:16], text)
	} else {
		// get strings except dynamic key
		b, e := Base64Decode(text[cKeyLen:])
		if e != nil {
			return "", e
		}
		text = string(b)
	}

	// text length
	textLen := len(text)
	if textLen <= 0 {
		err = fmt.Errorf("auth [%s] textLen <= 0", text)
		return
	}

	// keys
	box := arrayutils.Range(0, 256)
	//
	rndKey := make([]int, 0, 256)
	cryptKeyB := []byte(cryptKey)
	for i := 0; i < 256; i++ {
		pos := i % keyLen
		rndKey = append(rndKey, int(cryptKeyB[pos]))
	}

	j := 0
	for i := 0; i < 256; i++ {
		j = (j + box[i] + rndKey[i]) % 256
		box[i], box[j] = box[j], box[i]
	}

	textB := []byte(text)
	a := 0
	j = 0
	result := make([]byte, 0, textLen)
	for i := 0; i < textLen; i++ {
		a = (a + 1) % 256
		j = (j + box[a]) % 256
		box[a], box[j] = box[j], box[a]
		result = append(result, byte(int(textB[i])^(box[(box[a]+box[j])%256])))
	}

	if isEncode == ENCODE {
		// trim equal
		return keyC + Base64Encode(result), nil
	}

	// check expire time
	d, e := strconv.ParseInt(string(result[0:10]), 10, 0)
	if e != nil {
		err = fmt.Errorf("expires time error: %s", e.Error())
		return
	}

	if (d == 0 || d-timestamp > 0) && string(result[10:26]) == Md5Sum(string(result[26:]) + keyB)[0:16] {
		return string(result[26:]), nil
	}

	err = fmt.Errorf("Authcode text [%s] error", text)
	return
}
