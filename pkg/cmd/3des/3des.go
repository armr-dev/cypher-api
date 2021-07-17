package TripleDES

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

var key = []byte("123456781234567812345678")
var iv = []byte("43218765")

// 3DES encryption
func Encrypt(origData string) (string, error) {
	parsedData := []byte(origData)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	parsedData = PKCS5Padding(parsedData, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(parsedData))
	blockMode.CryptBlocks(encrypted, parsedData)

	return hex.EncodeToString(encrypted), nil
}

// 3DES decryption
func Decrypt(encrypted string) (string, error) {
	parsedData, _ := hex.DecodeString(encrypted)
	block, err := des.NewTripleDESCipher(key)
	
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(parsedData))

	blockMode.CryptBlocks(origData, parsedData)
	origData = PKCS5UnPadding(origData)

	return string(origData), nil
}


func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// remove the last byte unpadding times
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
