package strategy

import (
	"database/sql"
	"fmt"
	c "github.com/event-crawler/config"
	_ "github.com/lib/pq"
	"github.com/stvp/go-toml-config"

	"testing"
)

func TestFbCallPageApi(t *testing.T) {
	err := config.Parse(DEV_CONFIG_PATH)
	if err != nil {
		t.Fatalf("Fail on parsing config file %s", err)
	}
	db, err := sql.Open("postgres", *c.DbConStr)
	msgs := make(chan string)
	quit := make(chan string)
	go FbCallPageApi(msgs, quit, db, "8087014348")
	for {
		select {
		case s := <-msgs:
			fmt.Printf("'%s'\n", s)
		case err := <-quit:
			if err != "" {
				t.Errorf("Error on FsGetWithBxLink strategy Error:", err)
			}
			fmt.Println("Done")
			return
		}
	}

}
