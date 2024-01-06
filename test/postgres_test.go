package test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zomgra/tracker/pkg/barcode"
	"github.com/zomgra/tracker/pkg/bloomfilter"
	"github.com/zomgra/tracker/pkg/config"
)

func createShipmentTest(t *testing.T, barcode string) {
	err := client.Insert(barcode)
	require.NoError(t, err)
}

func getExistShipmentTest(t *testing.T, barcode string) {
	exist, err := client.Exists(barcode)
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
	bloomConfig := config.BloomFilterConfig{}
	config.SetBloomConfig(&bloomConfig)
	barcodes := make(chan string, n)
	filter := bloomfilter.NewBloomFilterHelper(&bloomConfig)
	for i := 0; i < n; i++ {
		barcode := barcode.GenerateBarCode(2, 2)
		createShipmentTest(t, barcode)
		barcodes <- barcode
	}
	close(barcodes)
	filter.Inject(client.InjectDataTo)

	for i := 0; i < n; i++ {
		shipmentBarcodeBytes, _ := json.Marshal(<-barcodes)
		ok := filter.Check(shipmentBarcodeBytes)
		require.True(t, ok)
	}

}
