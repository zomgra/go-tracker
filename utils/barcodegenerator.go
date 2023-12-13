package utils

import (
	"math/rand"
	"time"
)

var AvailableChars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

const MaxLengthBarCode = 20

func GenerateBarCode(charInStart, charInEnd int) string {
	rand.Seed(time.Now().UnixNano())

	var result []rune
	for i := 0; i < charInStart; i++ {
		result = append(result, AvailableChars[rand.Intn(len(AvailableChars))])
	}

	for i := 0; i < MaxLengthBarCode-(charInStart+charInEnd); i++ {
		result = append(result, rune(rand.Intn(10)+'0'))
	}

	for i := 0; i < charInEnd; i++ {
		result = append(result, AvailableChars[rand.Intn(len(AvailableChars))])
	}

	return string(result)
}
