package googlecse

import (
	"fmt"
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
	testSearchBxLinks(t)

}

func testSearchBxLinks(t *testing.T) {

	result, err := SearchBxLinks("bostancı gösteri merkezi")
	fmt.Println(result)
	if err != nil {
		t.Fatalf("Fail on exploring foursquare test venue, error: %s", err)
	}

	if len(result.Items) == 0 {
		t.Errorf("No error has occurred but SearchResult doesn't contain any items.", len(result.Items))
	}
}
