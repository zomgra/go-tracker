package db

import (
	"encoding/json"
	"testing"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/stretchr/testify/require"
	"github.com/zomgra/tracker/utils"
)

func createShipmentTest(t *testing.T, barcode string) {
	err := InsertShipment(barcode)
	require.NoError(t, err)
}

func getExistShipmentTest(t *testing.T, barcode string) {
	exist, err := ExistShipment(barcode)
	require.NoError(t, err)
	require.True(t, exist)
}

func TestIncludingShipment(t *testing.T) {
	barcode := utils.GenerateBarCode(2, 2)

	// create shipment
	createShipmentTest(t, barcode)

	// check shipment
	getExistShipmentTest(t, barcode)
}

func TestInjectingFromDb(t *testing.T) {
	n := 5
	barcodes := make(chan string, n)
	filter := bloom.NewWithEstimates(1000000, 0.01)

	for i := 0; i < n; i++ {
		barcode := utils.GenerateBarCode(2, 2)
		createShipmentTest(t, barcode)
		barcodes <- barcode
	}

	InjectDataTo(filter)

	for i := 0; i < n; i++ {
		shipmentBarcodeBytes, _ := json.Marshal(<-barcodes)
		ok := filter.Test(shipmentBarcodeBytes)
		require.True(t, ok)
	}

}
