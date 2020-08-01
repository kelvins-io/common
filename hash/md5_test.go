package hash

import (
	"fmt"
	"gitee.com/kelvins-io/common/convert"
	"testing"
)

func TestMD5Encode(t *testing.T) {
	fmt.Println(convert.Byte2Str(MD5Encode("hello world")))
}

func TestMD5EncodeToString(t *testing.T) {
	fmt.Println(MD5EncodeToString("hello world"))
}
