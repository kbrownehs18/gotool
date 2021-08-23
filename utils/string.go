package utils

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Reverse string reverse
func Reverse(s string) string {
	b := []byte(s)
	n := ""
	for i := len(b); i > 0; i-- {
		n += string(b[i-1])
	}
	return string(n)
}

// IsIP ip address is valid
func IsIP(ip string) bool {
	ips := strings.Split(ip, ".")
	if len(ips) != 4 {
		return false
	}
	for _, v := range ips {
		i, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if i < 0 || i > 255 {
			return false
		}
	}

	return true
}

// IsMac mac address is valid
func IsMac(mac string) bool {
	if len(mac) != 17 {
		return false
	}
	mac = strings.ToLower(mac)

	r := `^(?i:[0-9a-f]{1})(?i:[02468ace]{1}):(?i:[0-9a-f]{2}):(?i:[0-9a-f]{2}):(?i:[0-9a-f]{2}):(?i:[0-9a-f]{2}):(?i:[0-9a-f]{2})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return false
	}

	return reg.FindStringSubmatch(mac) == nil
}

// IsEmpty true: nil, "", false, 0, 0.0, {}, []
func IsEmpty(val interface{}) (b bool) {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.Bool:
		b = !val.(bool)
	case reflect.String:
		b = (len(val.(string)) == 0)
	case reflect.Array, reflect.Slice, reflect.Map:
		b = (v.Len() == 0)
	default:
		b = (v.Interface() == reflect.ValueOf(0).Interface() || v.Interface() == reflect.ValueOf(0.0).Interface())
	}

	return b
}

// IsBlank trim string then check
func IsBlank(str string) bool {
	return Trim(str) == ""
}

// Trim remove "", \r, \t, \n
func Trim(str string) string {
	return strings.Trim(str, " \r\n\t")
}

// Split split by match
func Split(str, match string) []string {
	re := regexp.MustCompile(match)
	return re.Split(str, -1)
}

// SplitBySpaceTab splite by space or tab
func SplitBySpaceTab(str string) []string {
	return Split(str, `[ \t]+`)
}

// IsMobile check chinese mobile number
func IsMobile(mobile string) bool {
	reg := `^1([3456789][0-9])\d{8}$`
	rgx := regexp.MustCompile(reg)

	return rgx.MatchString(mobile)
}
