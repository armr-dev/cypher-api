package Idea

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"

	"github.com/dgryski/go-idea"
)

var key = []byte("1234567890123456")
var iv = []byte("65432109")
var ideaBlockSize = 8

func Encrypt(data string) (string, error) {
	parsedData := []byte(data)
	block, err := idea.NewCipher(key)
	
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	parsedData = pkcs5Padding(parsedData, block.BlockSize())
	cryptText := make([]byte, len(parsedData))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cryptText, parsedData)

	return hex.EncodeToString(cryptText), nil
}

func Decrypt(data string) (string, error) {
	parsedData, _ := hex.DecodeString(data)
	block, err := idea.NewCipher(key)

	if err != nil {
		return "", err
	}

	if len(parsedData) < ideaBlockSize {
		panic("ciphertext too short")
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	cryptText := make([]byte, len(parsedData))
	blockMode.CryptBlocks(cryptText, parsedData)
	cryptText = pkcs5Depadding(cryptText, ideaBlockSize)

	return string(cryptText), nil
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcs5Depadding(cipherText []byte, blockSize int) []byte {
	padding := cipherText[len(cipherText)-1]
	out := bytes.Split(cipherText, []byte{padding})
	return out[0][:]
}
