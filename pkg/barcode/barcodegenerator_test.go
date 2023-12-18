package barcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBarcodeGeneratorMustBe20LengthWith4Chars(t *testing.T) {
	barcode := GenerateBarCode(2, 2)

	require.NotEmpty(t, barcode)
	require.Len(t, barcode, 20)
}

func TestBarcodeGeneratorMustBe20LengthWith5Chars(t *testing.T) {
	barcode := GenerateBarCode(2, 3)

	require.NotEmpty(t, barcode)
	require.Len(t, barcode, 20)
}
