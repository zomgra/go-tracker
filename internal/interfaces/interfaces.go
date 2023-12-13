package interfaces

import "github.com/zomgra/tracker/internal/domain"

type Repository interface {
	OnLoad() bool
	AddShipment(domain.Shipment)
	CheckShipment(barcode string) bool
	InjectFromDB()
}
