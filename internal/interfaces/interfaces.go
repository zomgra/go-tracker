package interfaces

import "github.com/zomgra/tracker/internal/models"

type Repository interface {
	OnLoad() bool
	AddShipment(models.Shipment)
	CheckShipment(barcode string) bool
	InjectFromDB()
}
