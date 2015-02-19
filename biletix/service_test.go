package biletix

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
	/*	testCrawlVenue(t)
		testCrawlVenueOfWrongUrl(t)*/
	testCrawlEvent(t)
	/*	testCrawlEventOfWrongUrl(t)*/

}

func testCrawlVenue(t *testing.T) {
	bxvenue, err := CrawlVenue("http://www.biletix.com/mekan/FU/ISTANBUL/tr")
	if err != nil {
		t.Fatalf("Fail on crawling biletix data, Error: %s", err)
	}
	if bxvenue.Name == "" {
		t.Errorf("No error has occurred but venue data has some problems, something must be wrong!")
	}
}

func testCrawlVenueOfWrongUrl(t *testing.T) {
	_, err := CrawlVenue("http://www.biletix.com/thisisawrongurl")
	if err == nil {
		t.Errorf("Expected error but didn't occur", err)
	}
}

func testCrawlEvent(t *testing.T) {
	bxevent, err := CrawlEvent("http://www.biletix.com/etkinlik/SKPO2/TURKIYE/tr")
	if err != nil {
		t.Fatalf("Fail on crawling biletix data, Error: %s", err)
	}

	if bxevent.Name == "" {
		t.Errorf("No error has occurred but venue data has some problems, something must be wrong!")
	}
}
func testCrawlEventOfWrongUrl(t *testing.T) {
	_, err := CrawlEvent("http://www.biletix.com/thisisawrongurl")
	if err == nil {
		t.Errorf("Expected error but didn't occur")
	}
}
