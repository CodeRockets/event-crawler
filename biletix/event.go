package biletix

import (
	"time"
)

type BxEvent struct {
	Id             int
	Name           string
	BxLink         string
	BxId           string
	VenueId        int
	EventImage     string
	EventDesc      string
	EventPrice     string
	EventDateStart time.Time
	PurchaseLink   string
}
