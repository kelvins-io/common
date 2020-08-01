package t_encrypt

import (
	"encoding/base64"
	"encoding/hex"
	"gitee.com/kelvins-io/common/crypt"
)

const trackSalt = "ffefff"
const trackIv = "U335YPjWDpaESQkb"

// encrypt
func TrackEncrypt(key string, str string) (string, error) {
	if CheckTrackEncryption(str) {
		return str, nil
	}

	crypted, err := crypt.AESEncrypt([]byte(str), []byte(key), []byte(trackIv))
	if err != nil {
		return "", err
	}

	return genSalt() + base64.StdEncoding.EncodeToString(crypted), nil
}

// decrypt
func TrackDecrypt(key string, str string) (string, error) {
	if !CheckTrackEncryption(str) {
		return str, nil
	}
	// 移除salt前缀
	s := str[4:]
	// base64解析
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	decode, err := crypt.AESDecrypt(data, []byte(key), []byte(trackIv))
	if err != nil {
		return "", err
	}

	return string(decode), nil
}

// ecb方式aes加密
func TrackECBEncrypt(key string, str string) (string, error) {
	if CheckTrackEncryption(str) {
		return str, nil
	}

	return genSalt() + base64.StdEncoding.EncodeToString(crypt.AesECBEncrypt([]byte(str), []byte(key))), nil
}

// ecb方式aes解密
func TrackECBDecrypt(key string, str string) (string, error) {
	if !CheckTrackEncryption(str) {
		return str, nil
	}
	// 移除salt前缀
	s := str[4:]
	// base64解析
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(crypt.AesECBDecrypt(data, []byte(key))), nil
}

// check encryption
func CheckTrackEncryption(str string) bool {
	if len(str) <= 4 {
		return false
	}

	return str[:4] == genSalt()
}

func genSalt() string {
	b, _ := hex.DecodeString(trackSalt)
	return base64.StdEncoding.EncodeToString(b)
}
