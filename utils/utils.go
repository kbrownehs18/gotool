package utils

import (
	"reflect"
)

// Contains item is in map/slice/array
func Contains(haystack interface{}, needle interface{}) bool {
	targetValue := reflect.ValueOf(haystack)
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == needle {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(needle)).IsValid() {
			return true
		}
	}

	return false
}
