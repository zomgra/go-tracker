package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zomgra/tracker/pkg/barcode"
)

func TestBarcodeGeneratorMustBe20LengthWith4Chars(t *testing.T) {
	barcode := barcode.GenerateBarCode(2, 2)

	require.NotEmpty(t, barcode)
	require.Len(t, barcode, 20)
}

func TestBarcodeGeneratorMustBe20LengthWith5Chars(t *testing.T) {
	barcode := barcode.GenerateBarCode(2, 3)

	require.NotEmpty(t, barcode)
	require.Len(t, barcode, 20)
}
