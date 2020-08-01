package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// SHA1 SHA1哈希加密
func SHA1(plainText []byte) string {
	sha := sha1.New()
	sha.Write(plainText)
	return hex.EncodeToString(sha.Sum(nil))
}

// MD5 MD5哈希加密， 返回32位字符串
func MD5(plainText []byte) string {
	m := md5.New()
	m.Write(plainText)
	return hex.EncodeToString(m.Sum(nil))
}

// AESEncrypt AES加密
func AESEncrypt(plaintext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext = pkcsPadding(plaintext, block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	mode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

// AESDecrypt AES 解密
func AESDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	mode.CryptBlocks(origData, crypted)
	origData = pkcsUnPadding(origData)
	return origData, nil
}

func pkcsPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcsUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesECBEncrypt(plaintext []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key[:aes.BlockSize])
	if err != nil {
		panic(err.Error())
	}

	plaintext = zeroPadding(plaintext, aes.BlockSize)

	ciphertext := make([]byte, 0)
	text := make([]byte, 16)
	for len(plaintext) > 0 {
		cipher.Encrypt(text, plaintext)
		plaintext = plaintext[aes.BlockSize:]
		ciphertext = append(ciphertext, text...)
	}
	return ciphertext
}

func AesECBDecrypt(ciphertext []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key[:aes.BlockSize])
	if err != nil {
		panic(err.Error())
	}

	plaintext := make([]byte, 0)
	text := make([]byte, 16)
	for len(ciphertext) > 0 {
		cipher.Decrypt(text, ciphertext)
		ciphertext = ciphertext[aes.BlockSize:]
		plaintext = append(plaintext, text...)
	}

	plaintext = zeroUnPadding(plaintext)

	return plaintext
}

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(data []byte) []byte {
	data = bytes.TrimFunc(data, func(r rune) bool {
		return r == rune(0)
	})

	return data
}