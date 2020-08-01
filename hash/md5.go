package hash

import (
	"crypto/md5"
	"encoding/hex"
	"gitee.com/kelvins-io/common/convert"
)

func MD5Encode(data string) []byte {
	h := md5.New()
	h.Write(convert.Str2Byte(data))
	return h.Sum(nil)
}

func MD5EncodeToString(data string) string {
	y := make([]byte, 32)
	hex.Encode(y, MD5Encode(data))
	return convert.Byte2Str(y)
}
