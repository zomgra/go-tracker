package models

import (
	"math/rand"
	"time"
)

type Shipment struct {
	Barcode string
}

func (shipment *Shipment) GenerateShipment() {
	generator := BarCodeGenerator{}
	code := generator.GenerateBarCode(2, 2)
	shipment.Barcode = code
}

type BarCodeGenerator struct{}

var AvailableChars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

const MaxLengthBarCode = 20

func (g *BarCodeGenerator) GenerateBarCode(charInStart, charInEnd int) string {
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
