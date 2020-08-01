package t_encrypt

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	var str = "fdf"
	var key = "2f43&^#fd"

	crypt, _ := TrackEncrypt(key, str)
	fmt.Println(crypt)

	fmt.Println(TrackDecrypt(key, crypt))

}

func TestECBEncrypt(t *testing.T) {
	var str = "dffdf234"
	var key = "dfdf&^#d43s90G"

	crypt, _ := TrackECBEncrypt(key, str)
	fmt.Println(crypt)

	fmt.Println(TrackECBDecrypt(key, crypt))
}
