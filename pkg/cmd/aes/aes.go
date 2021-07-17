package AES

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)
var key = []byte("1234567890123456")
var iv = []byte("6543210987654321")

func Encrypt(data string) (string, error) {
	parsedData := []byte(data)
	block, err := aes.NewCipher(key)
	
	if err != nil {
		fmt.Println("AQUI", err)
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
	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	if len(parsedData) < aes.BlockSize {
		panic("ciphertext too short")
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	cryptText := make([]byte, len(parsedData))
	blockMode.CryptBlocks(cryptText, parsedData)
	cryptText = pkcs5Depadding(cryptText, aes.BlockSize)

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
