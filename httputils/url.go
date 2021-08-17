package httputils

import (
	"net/http"
)

// URLEncode urlencode
func URLEncode(params interface{}) string {
	q, ok := params.(string)
	if ok {
		return http.QueryEscape(q)
	}
	m, ok := params.(map[string]string)
	if ok {
		val := http.Values{}
		for k, v := range m {
			val.Set(k, v)
		}

		return val.Encode()
	}

	return ""
}

// URLDecode urldecode
func URLDecode(str string) (string, error) {
	return http.QueryUnescape(str)
}
