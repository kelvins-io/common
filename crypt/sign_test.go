package crypt

import (
	"testing"
	"fmt"
)

func TestMd5Sign(t *testing.T) {
	fmt.Println(Md5Sign(map[string]string{
		"timestamp":"12312313",
		"abcc": "kkk",
		"cccc": "zzz",
		"xxx": "312",
		"sign": "9ADCEAA7C1C1979x2214BECF7747E05DAA",
	}, "123123"))
}
