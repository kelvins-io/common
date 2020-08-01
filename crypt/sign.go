package crypt

import (
	"bytes"
	"sort"
	"strings"
)

func CheckMd5Sign(sign string, params map[string]string, apiKey string) bool {
	return sign == Md5Sign(params, apiKey)
}

func Md5Sign(params map[string]string, apiKey string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var bf bytes.Buffer
	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		bf.WriteString(k)
		bf.WriteByte('=')
		bf.WriteString(v)
		bf.WriteByte('&')
	}
	bf.WriteString("key=")
	bf.WriteString(apiKey)

	return strings.ToUpper(MD5(bf.Bytes()))
}

// for sms api md5 sign
func Md5SignSms(params map[string]string, apiKey string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var bf bytes.Buffer
	for i, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		bf.WriteString(k)
		bf.WriteByte('=')
		bf.WriteString(v)
		if i < len(keys)-1 {
			bf.WriteByte('&')
		}
	}

	bf.WriteString(apiKey)

	return MD5(bf.Bytes())
}
