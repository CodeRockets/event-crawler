package strategy

import (
	"database/sql"
	"fmt"
	fb "github.com/event-crawler/facebook"
	"time"
)

func FbCallPageApi(msgs chan<- string, quit chan<- string, db *sql.DB, fbId string) {
	fmt.Println("...")
	page := new(fb.Page)
	page.FbId = fbId

	msgs <- "Call page api of " + fbId + " from facebook"
	err := page.Call()
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}
	msgs <- "Page Info has collected (" + fbId + ")"
	pageno, err := page.Insert(db)
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}
	msgs <- "Page Info has inserted to db id:" + string(pageno)
	msgs <- "Call events api of page " + fbId + " from facebook"
	eventlist, err := page.GetEvents()
	if err != nil {
		quit <- fmt.Sprintf("%s", err)
	}
	for i := 0; i < len(eventlist); i++ {
		eventlist[i].Call()
		msgs <- "      " + eventlist[i].Name + " event info crawled"
		eventlist[i].PageId = fbId

		err = eventlist[i].CheckVenueIsInDb(db)
		if err != nil {
			quit <- fmt.Sprintf("%s", err)
		}

		_, err = eventlist[i].Insert(db)
		if err != nil {
			quit <- fmt.Sprintf("%s", err)
		}
		msgs <- "      " + eventlist[i].Name + " event info inserted into database"
		time.Sleep(time.Second * 1)
	}
	quit <- fmt.Sprintf("%s", "...DONE...")

}
