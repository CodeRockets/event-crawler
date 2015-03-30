package facebook

import (
	"database/sql"

	c "github.com/event-crawler/config"

	_ "github.com/lib/pq"
	"github.com/stvp/go-toml-config"
	"testing"
)

func TestEventCall(t *testing.T) {
	err := config.Parse(DEV_CONFIG_PATH)
	if err != nil {
		t.Fatalf("Fail on parsing config file %s", err)
	}
	db, _ := sql.Open("postgres", *c.DbConStr)
	event := new(Event)
	event.FbId = "383414275117138"
	err = event.Call()
	event.Insert(db)
	if err != nil {
		t.Fatalf("Fail on calling fb api %s", err)
	}
}
