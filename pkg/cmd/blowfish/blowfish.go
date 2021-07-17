package Blowfish

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math"
)

func SplitText(text []byte) ([]byte, []byte) {
	textLen := len(text)
	halfLen := int(math.Ceil(float64(textLen) / 2))

	xL := text[0:halfLen]
	xR := text[halfLen:textLen]

	return xL, xR
}

func MergeText(xL, xR uint32) []byte {
	text1 := make([]byte, 4)
	binary.BigEndian.PutUint32(text1, xL)
	text2 := make([]byte, 4)
	binary.BigEndian.PutUint32(text2, xR)
	// text := uint32ToByte(xL)
	text := append(text1, text2...)

	return text
}

func F(xL uint32) uint32 {
	convertedText := make([]byte, 4)
	binary.BigEndian.PutUint32(convertedText, xL)
	firstHalf, secondHalf := SplitText([]byte(convertedText))

	a, b := SplitText(firstHalf)
	c, d := SplitText(secondHalf)

	modOp := uint64(math.Pow(2, 32))
	op1 := ((uint64(sBox0[a[0]] + sBox1[b[0]])) % modOp) ^ uint64(sBox2[c[0]])
	op2 := uint64(sBox3[d[0]]) % modOp

	return uint32(op1 + op2)
}

func EncryptBlock(blockText []byte) []byte {
	var xL, xR uint32
	auxL, auxR := SplitText(blockText[:])

	xL = binary.BigEndian.Uint32(auxL)
	xR = binary.BigEndian.Uint32(auxR)

	var tmp uint32

	for i := 0; i < 16; i++ {
		xL = xL ^ pArray[i]
		xR = F(xL) ^ xR

		tmp = xL
		xL = xR
		xR = tmp
	}
	// Undo the swap from the last for iteration
	tmp = xL
	xL = xR
	xR = tmp

	xR = xR ^ pArray[16]
	xL = xL ^ pArray[17]

	cypheredText := MergeText(xL, xR)

	return cypheredText
}

func DecryptBlock(blockText []byte) []byte {
	var xL, xR uint32
	auxL, auxR := SplitText(blockText[:])

	xL = binary.BigEndian.Uint32(auxL)
	xR = binary.BigEndian.Uint32(auxR)

	var tmp uint32

	xL = xL ^ pArray[17]
	xR = xR ^ pArray[16]

	tmp = xL
	xL = xR
	xR = tmp

	for i := 15; i >= 0; i-- {
		tmp = xL
		xL = xR
		xR = tmp

		xR = F(xL) ^ xR
		xL = xL ^ pArray[i]
	}

	decypheredText := MergeText(xL, xR)

	return decypheredText
}

func Encrypt(text string) string {
	var cypheredText []byte
	clearText := []byte(text)

	// Add spacebars (32) if the text lenght is < 8
	nBlocks := len(clearText) / 8

	if len(clearText)%8 != 0 {
		for i := len(clearText); i < 8; i++ {
			clearText = append(clearText, 32)
		}

		nBlocks++
	}

	for i := 0; i < nBlocks; i++ {
		cypheredText = append(cypheredText, EncryptBlock(clearText[8*i:(i+1)*8])...)
	}

	return hex.EncodeToString(cypheredText)
}

func Decrypt(text string) (string, error) {
	var decypheredText []byte
	opaqueText, err := hex.DecodeString(text)

	if len(text)%8 != 0 {
		return "", errors.New("corrupted encrypted message")
	}

	// Add spacebars (32) if the text lenght is < 8
	nBlocks := len(opaqueText) / 8

	if len(opaqueText)%8 != 0 {
		for i := len(opaqueText); i < 8; i++ {
			opaqueText = append(opaqueText, 32)
		}

		nBlocks++
	}

	for i := 0; i < nBlocks; i++ {
		decypheredText = append(decypheredText, DecryptBlock(opaqueText[8*i:(i+1)*8])...)
	}

	return string(decypheredText), err
}
