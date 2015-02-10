package strategy

import (
	"database/sql"
	"fmt"
	bx "github.com/event-crawler/biletix"
	fs "github.com/event-crawler/foursquare"
	cs "github.com/event-crawler/googlecse"
)

func FsGetWithBxLink(msgs chan<- string, quit chan<- string, db *sql.DB, fsId string, bxLink string) {

	fsGetRes, err := fs.Get(fsId)
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}

	fsvenue := fsGetRes.Response.MVenue
	msgs <- fsvenue.Name + " venue info has collected from foursquare"
	lastid, err := fs.InsVenue(db, fsvenue)
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}
	msgs <- fsvenue.Name + " inserted into database"
	_, err = bx.UpdateVenueBxUrlByFsId(db, fsId, bxLink)
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}
	msgs <- fsvenue.Name + " biletix link updated manually"
	bxvenue, err := bx.CrawlVenue(bxLink)
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}
	msgs <- fsvenue.Name + " venue info has collected from biletix"
	bxvenue.VenueId = lastid
	_, err = bx.UpdateVenue(db, bxvenue)
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}
	msgs <- fsvenue.Name + " venue info updated"
	for i := 0; i < len(bxvenue.BxEventLinks); i++ {
		bxevent, _ := bx.CrawlEvent(bxvenue.BxEventLinks[i])
		msgs <- "      " + bxevent.Name + " event info crawled"
		bxevent.VenueId = lastid
		bx.InsertEvent(db, bxevent)
		msgs <- "      " + bxevent.Name + " event info inserted into database"
	}
	quit <- fmt.Sprintf("%s", "...DONE...")
}

func FsGetWithGCSE(msgs chan<- string, quit chan<- string, db *sql.DB, fsId string) {

	fsGetRes, err := fs.Get(fsId)
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}

	fsvenue := fsGetRes.Response.MVenue
	msgs <- fsvenue.Name + " venue info has collected from foursquare"
	lastid, err := fs.InsVenue(db, fsvenue)
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}
	msgs <- fsvenue.Name + " inserted into database"
	cseresult, err := cs.SearchBxLinks(fsvenue.Name)
	bxLink := cseresult.Items[0].Link
	msgs <- fsvenue.Name + " biletix link found by Google Custom Search"
	_, err = bx.UpdateVenueBxUrlByFsId(db, fsId, bxLink)
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}
	msgs <- fsvenue.Name + " biletix link updated"
	bxvenue, err := bx.CrawlVenue(bxLink)
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}
	msgs <- fsvenue.Name + " venue info has collected from biletix"
	bxvenue.VenueId = lastid
	_, err = bx.UpdateVenue(db, bxvenue)
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}
	msgs <- fsvenue.Name + " venue info updated"
	for i := 0; i < len(bxvenue.BxEventLinks); i++ {
		bxevent, _ := bx.CrawlEvent(bxvenue.BxEventLinks[i])
		msgs <- "      " + bxevent.Name + " event info crawled"
		bxevent.VenueId = lastid
		bx.InsertEvent(db, bxevent)
		msgs <- "      " + bxevent.Name + " event info inserted into database"
	}
	quit <- fmt.Sprintf("%s", "...DONE...")
}
