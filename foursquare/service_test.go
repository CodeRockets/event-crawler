package foursquare

import (
	"github.com/stvp/go-toml-config"
	"testing"
)

const (
	DEV_CONFIG_PATH  = "../config/dev.conf"
	PROD_CONFIG_PATH = "../config/prod.conf"
)

func TestService(t *testing.T) {
	err := config.Parse(DEV_CONFIG_PATH)
	if err != nil {
		t.Fatalf("Fail on parsing config file %s", err)
	}
	testExplore(t)
	testGet(t)
	testGetWithWrongId(t)
}

func testExplore(t *testing.T) {

	fsVenueList, err := Explore("İstanbul", "music")
	if err != nil {
		t.Fatalf("Fail on exploring foursquare test venue, error: %s", err)
	}
	if len(fsVenueList.Response.Groups[0].Items) == 0 {
		t.Errorf("No error has occurred but FsVenueList doesn't contain any items.", len(fsVenueList.Response.Groups[0].Items))
	}
}

func testGet(t *testing.T) {
	_, err := Get("4b4ef332f964a52091f726e3")
	if err != nil {
		t.Errorf("Fail on getting foursquare test venue by id, error:  %s", err)
	}
}
func testGetWithWrongId(t *testing.T) {

	fsVenue, _ := Get("thisisawrongid")
	if len(fsVenue.Response.MVenue.Id) != 0 {
		t.Errorf("Empty string was expected as id but it is not, id: %s", fsVenue.Response.MVenue.Id)
	}
}
