package foursquare

import (
	"database/sql"
	c "github.com/event-crawler/config"
	_ "github.com/lib/pq"
	"github.com/stvp/go-toml-config"
	"testing"
)

func TestDataStore(t *testing.T) {
	err := config.Parse(DEV_CONFIG_PATH)
	if err != nil {
		t.Fatalf("Fail on parsing config file %s", err)
	}

	db, _ := sql.Open("postgres", *c.DbConStr)
	venue := new(FsVenue)
	venue.Id = "testId"
	venue.Location.FormattedAddress = []string{"test", "test1"}
	lastid, err := InsVenue(db, *venue)

	if err != nil {
		t.Fatalf("Fail on inserting test venue, error: %s", err)
	}
	if lastid == 0 {
		t.Fatal("No error has occurred but last inserted id is zero")
	}

	result, err := DelByFsId(db, venue.Id)
	if err != nil {
		t.Fatalf("Fail on deleting test venue, error: %s", err)
	}

	rowseffected, err := result.RowsAffected()
	if rowseffected == 0 {
		t.Errorf("No error has occurred and no row has effected, effected row count %s", rowseffected)
	}

}
