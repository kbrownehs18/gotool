package httputils

import "net/url"

// URLEncode urlencode
func URLEncode(params interface{}) string {
	q, ok := params.(string)
	if ok {
		return url.QueryEscape(q)
	}
	m, ok := params.(map[string]string)
	if ok {
		val := url.Values{}
		for k, v := range m {
			val.Set(k, v)
		}

		return val.Encode()
	}

	return ""
}

// URLDecode urldecode
func URLDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}
