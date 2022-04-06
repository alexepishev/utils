package utils

import (
	"fmt"
	"strings"
)

func MapGetKeyValues(s string) map[string]string {

	env := GetEnvironment()
	pairs := strings.Split(s, ",")

	m := make(map[string]string)

	for _, p := range pairs {

		if IsEmpty(p) {
			continue
		}
		kv := strings.SplitN(p, "=", 2)
		k := strings.TrimSpace(kv[0])
		if len(kv) > 1 {
			v := strings.TrimSpace(kv[1])
			if strings.HasPrefix(v, "${") && strings.HasSuffix(v, "}") {
				ed := strings.SplitN(v[2:len(v)-1], ":", 2)
				e, d := ed[0], ed[1]
				v = env.Get(e, "").(string)
				if v == "" && d != "" {
					v = d
				}
			}
			m[k] = v
		} else {
			m[k] = ""
		}
	}
	return m
}

func MapToArrayWithSeparator(m map[string]string, s string) []string {

	var arr []string
	if m == nil {
		return arr
	}
	for k, v := range m {
		if IsEmpty(v) {
			arr = append(arr, k)
		} else {
			arr = append(arr, fmt.Sprintf("%s%s%v", k, s, v))
		}
	}
	return arr
}

func MapToArray(m map[string]string) []string {
	return MapToArrayWithSeparator(m, "=")
}
