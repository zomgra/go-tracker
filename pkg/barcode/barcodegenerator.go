package barcode

import (
	"math/rand"
	"time"
)

var alphabetRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

const (
	MaxLengthBarCode = 20
	//AvailableChars   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func GenerateBarCode(charInStart, charInEnd int) string {
	rand.Seed(time.Now().UnixNano())

	var result []rune
	for i := 0; i < charInStart; i++ {
		result = append(result, alphabetRunes[rand.Intn(len(alphabetRunes))])
	}

	for i := 0; i < MaxLengthBarCode-(charInStart+charInEnd); i++ {
		result = append(result, rune(rand.Intn(10)+'0'))
	}

	for i := 0; i < charInEnd; i++ {
		result = append(result, alphabetRunes[rand.Intn(len(alphabetRunes))])
	}

	return string(result)
}
