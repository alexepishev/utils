package utils

import (
	"io/ioutil"
	"reflect"
	"strings"
)

// IsEmpty ...
func IsEmpty(v interface{}) bool {

	switch v.(type) {
	case string:
		return len(strings.TrimSpace(v.(string))) == 0
	case int:
		return v.(int) == 0
	case bool:
		return v.(bool)
	default:
		return v == nil
	}
}

// Contains ...
func Contains(items interface{}, item interface{}) bool {

	arrV := reflect.ValueOf(items)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == item {
				return true
			}
		}
	}
	return false
}

func Content(contentOrPath string) ([]byte, error) {

	if FileExists(contentOrPath) {
		b, err := ioutil.ReadFile(contentOrPath)
		if err != nil {
			return nil, err
		}
		return b, nil
	}
	return []byte(contentOrPath), nil
}
