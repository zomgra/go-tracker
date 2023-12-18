package test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zomgra/tracker/pkg/barcode"
	"github.com/zomgra/tracker/pkg/bloomfilter"
)

func createShipmentTest(t *testing.T, barcode string) {
	err := client.Insert(barcode)
	require.NoError(t, err)
}

func getExistShipmentTest(t *testing.T, barcode string) {
	exist, err := client.Exist(barcode)
	require.NoError(t, err)
	require.True(t, exist)
}

func TestIncludingShipment(t *testing.T) {
	barcode := barcode.GenerateBarCode(2, 2)

	// create shipment
	createShipmentTest(t, barcode)

	// check shipment
	getExistShipmentTest(t, barcode)
}

func TestInjectingFromDb(t *testing.T) {
	n := 5
	barcodes := make(chan string, n)
	filter := bloomfilter.NewBloomFilterHelper()
	for i := 0; i < n; i++ {
		barcode := barcode.GenerateBarCode(2, 2)
		createShipmentTest(t, barcode)
		barcodes <- barcode
	}

	client.InjectDataTo(filter)

	for i := 0; i < n; i++ {
		shipmentBarcodeBytes, _ := json.Marshal(<-barcodes)
		ok := filter.Check(shipmentBarcodeBytes)
		require.True(t, ok)
	}

}
