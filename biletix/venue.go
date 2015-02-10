package biletix

import (
	"time"
)

type BxVenue struct {
	VenueId      int
	Name         string
	Image        string
	Desc         string
	Directions   string
	UpdateDate   time.Time
	Status       int
	Tags         string
	BxUrl        string
	BxEventLinks []string
}
