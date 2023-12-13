package interfaces

import "github.com/zomgra/tracker/internal/models"

type Repository interface {
	AddShipment(models.Shipment)
	CheckShipment(barcode string) bool
	InjectFromDB()
}
