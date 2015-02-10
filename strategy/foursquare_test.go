package strategy

import (
	"database/sql"
	"fmt"
	c "github.com/event-crawler/config"
	_ "github.com/lib/pq"
	"github.com/stvp/go-toml-config"

	"testing"
)

const (
	DEV_CONFIG_PATH  = "../config/dev.conf"
	PROD_CONFIG_PATH = "../config/prod.conf"
)

func TestFsGetWithBxLink(t *testing.T) {
	err := config.Parse(DEV_CONFIG_PATH)
	if err != nil {
		t.Fatalf("Fail on parsing config file %s", err)
	}
	db, err := sql.Open("postgres", *c.DbConStr)
	msgs := make(chan string)
	quit := make(chan struct{ error })
	go FsGetWithBxLink(msgs, quit, db, "52851e6f498e516d31b588d1", "http://www.biletix.com/mekan/09/TURKIYE/tr")
	for {
		select {
		case s := <-msgs:
			fmt.Printf("'%s'\n", s)
		case err := <-quit:
			if err.error != nil {
				t.Errorf("Error on FsGetWithBxLink strategy Error:", err)
			}
			fmt.Println("Done")
			return

		}
	}

}

func TestFsGetWithGoogleCSE(t *testing.T) {
	err := config.Parse(DEV_CONFIG_PATH)
	if err != nil {
		t.Fatalf("Fail on parsing config file %s", err)
	}
	db, err := sql.Open("postgres", *c.DbConStr)
	msgs := make(chan string)
	quit := make(chan struct{ error })
	go FsGetWithGCSE(msgs, quit, db, "52851e6f498e516d31b588d1")
	for {
		select {
		case s := <-msgs:
			fmt.Printf("'%s'\n", s)
		case err := <-quit:
			if err.error != nil {
				t.Errorf("Error on FsGetWithBxLink strategy Error:", err)
			}
			fmt.Println("Done")
			return

		}
	}
}
