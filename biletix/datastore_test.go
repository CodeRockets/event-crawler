package biletix

import (
	"database/sql"
	c "github.com/event-crawler/config"
	_ "github.com/lib/pq"
	"github.com/stvp/go-toml-config"
	"testing"
	"time"
)

func TestDatastore(t *testing.T) {
	err := config.Parse(DEV_CONFIG_PATH)
	if err != nil {
		t.Fatalf("Fail on parsing config file %s", err)
	}

	db, _ := sql.Open("postgres", *c.DbConStr)
	testUpdateVenue(t, db)
	testUpdateVenueBxUrlByFsId(t, db)
	testInsertEvent(t, db)
}

func testUpdateVenue(t *testing.T, db *sql.DB) {

	bxVenue := BxVenue{
		VenueId:    148,
		Name:       "test-name",
		Image:      "test-image",
		Desc:       "test-description",
		Directions: "test-directions",
		UpdateDate: time.Now(),
		Status:     109, //109 Test Venue
		Tags:       "tags,test",
	}

	updateresult, err := UpdateVenue(db, bxVenue)
	if err != nil {
		t.Errorf("Fail on updating test venue, error: %s", err)
	}
	rowseffected, _ := updateresult.RowsAffected()
	if rowseffected == 0 {
		t.Errorf("No error has occurred and no row has effected, effected row count is ", rowseffected)
	}

}

func testUpdateVenueBxUrlByFsId(t *testing.T, db *sql.DB) {

	updateresult, err := UpdateVenueBxUrlByFsId(db, "506696f9e4b08e8fd358c078", "test-url")
	if err != nil {
		t.Errorf("Fail on updating bx_url of test venue by fs_id , error: %s", err)
	}
	rowseffected, _ := updateresult.RowsAffected()
	if rowseffected == 0 {
		t.Errorf("No error has occurred and no row has effected, effected row count is ", rowseffected)
	}

}

func testInsertEvent(t *testing.T, db *sql.DB) {

	bxEvent := BxEvent{
		EventDesc:      "test-desc",
		BxId:           "test-bxid",
		BxLink:         "test-bxlink",
		EventDateStart: time.Now(),
		EventImage:     "test-eventimage",
		PurchaseLink:   "test-purchaseLink",
		EventPrice:     "test-eventprice",
		Name:           "test-eventName",
		VenueId:        148,
		PriceList:      "{10.22,11.12}",
	}

	result, err := InsertEvent(db, bxEvent)

	if err != nil {
		t.Fatalf("Fail on updating test event, error: %s", err)
	}
	rowseffected, err := result.RowsAffected()
	if rowseffected == 0 {
		t.Errorf("No error has occurred and no row has effected, effected row count %s", rowseffected)
	}

	result, err = DelEventByBxId(db, bxEvent.BxId)
	if err != nil {
		t.Fatalf("Fail on deleting test event, error: %s", err)
	}

	rowseffected, err = result.RowsAffected()
	if rowseffected == 0 {
		t.Errorf("No error has occurred and no row has effected, effected row count %s", rowseffected)
	}

}
