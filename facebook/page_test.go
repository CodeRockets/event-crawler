package facebook

import (
	"database/sql"
	"fmt"
	c "github.com/event-crawler/config"

	_ "github.com/lib/pq"
	"github.com/stvp/go-toml-config"
	"testing"
)

func TestInsert(t *testing.T) {
	err := config.Parse(DEV_CONFIG_PATH)
	if err != nil {
		t.Fatalf("Fail on parsing config file %s", err)
	}

	db, _ := sql.Open("postgres", *c.DbConStr)
	page := new(Page)
	page.FbId = "testId"
	page.Name = "TestName"
	page.Category = "testcategory"
	page.Checkins = 1000
	page.Cover.Source = "testimage"
	page.Description = "testdesc"
	page.Likes = 1000
	page.Location.Latitude = 49.565
	page.Location.Longitude = 49.565
	page.Location.Address = "testadd"
	page.Location.City = "testcity"
	page.UserName = "testusername"
	page.IsVenue = true

	lastid, err := page.Insert(db)
	if err != nil {
		t.Fatalf("Fail on inserting test page, error: %s", err)
	}
	if lastid == 0 {
		t.Fatal("No error has occurred but last inserted id is zero")
	}
	_, err = page.Delete(db, "testId")
	if err != nil {
		t.Fatalf("Fail on deleting test page, error: %s", err)
	}

}

func TestPageCall(t *testing.T) {
	err := config.Parse(DEV_CONFIG_PATH)
	if err != nil {
		t.Fatalf("Fail on parsing config file %s", err)
	}

	page := new(Page)
	page.FbId = "8087014348"
	err = page.Call()

	if err != nil {
		t.Fatalf("Fail on calling fb api %s", err)
	}
}

func TestGetEvents(t *testing.T) {
	err := config.Parse(DEV_CONFIG_PATH)
	if err != nil {
		t.Fatalf("Fail on parsing config file %s", err)
	}

	page := new(Page)
	page.FbId = "8087014348"
	_, err = page.GetEvents()
	fmt.Println("...")

	if err != nil {
		t.Fatalf("Fail on calling fb api %s", err)
	}
}
