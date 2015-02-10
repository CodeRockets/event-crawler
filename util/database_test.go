package util

import (
	"database/sql"
	c "github.com/event-crawler/config"
	_ "github.com/lib/pq"
	"github.com/stvp/go-toml-config"
	"testing"
)

const (
	DEV_CONFIG_PATH  = "../config/dev.conf"
	PROD_CONFIG_PATH = "../config/prod.conf"
)

func TestEventDatabase(t *testing.T) {
	err := config.Parse(DEV_CONFIG_PATH)
	if err != nil {
		t.Fatalf("Fail on parsing config file %s", err)
	}
	db, _ := sql.Open("postgres", *c.DbConStr)

	err = CreateTables(db)
	if err != nil {
		t.Fatalf("Fail on creating tables. Err:", err)
	}
	err = DropTables(db)
	if err != nil {
		t.Fatalf("Fail on dropping tables. Err:", err)
	}

}
