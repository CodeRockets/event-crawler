package util

import (
	"database/sql"
)

func CreateTables(db *sql.DB) error {
	_, err := db.Exec(createEventSql)
	if err != nil {
		return err
	}
	_, err = db.Exec(createVenueSql)
	return nil
}
func DropTables(db *sql.DB) error {
	_, err := db.Exec(dropEventSql)
	if err != nil {
		return err
	}
	_, err = db.Exec(dropVenueSql)
	return err
}

const createEventSql = `CREATE TABLE "public"."event" (
"id" int4 DEFAULT nextval('event_id_seq'::regclass) NOT NULL,
"venue_id" int4,
"event_name" varchar(255) COLLATE "default",
"bx_event_link" varchar(255) COLLATE "default",
"created_at" timestamp(6),
"updated_at" timestamp(6),
"status" int2,
"event_image" varchar(255) COLLATE "default",
"bx_event_id" varchar(50) COLLATE "default",
"event_desc" text COLLATE "default",
"event_date_start" timestamp(6),
"event_price" text COLLATE "default",
"purchase_link" varchar(255) COLLATE "default",
CONSTRAINT "event_pkey" PRIMARY KEY ("id")
)
WITH (OIDS=FALSE)
;`

const dropEventSql = "drop table event"

const createVenueSql = `CREATE TABLE "public"."venue" (
"id" int4 DEFAULT nextval('venue_id_seq'::regclass) NOT NULL,
"fs_id" varchar(50) COLLATE "default",
"name" varchar(255) COLLATE "default",
"phone" varchar(50) COLLATE "default",
"twitter" varchar(50) COLLATE "default",
"facebook_id" varchar(50) COLLATE "default",
"facebook_username" varchar(100) COLLATE "default",
"url" varchar(100) COLLATE "default",
"city" varchar(50) COLLATE "default",
"country" varchar(50) COLLATE "default",
"formatted_address" varchar(500) COLLATE "default",
"fs_tags" text[] COLLATE "default",
"fs_rating" float4,
"fs_rating_signals" int4,
"address" varchar(500) COLLATE "default",
"created_at" timestamp(6),
"updated_at" timestamp(6),
"lat" numeric(9,6),
"lon" numeric(9,6),
"bx_tag" varchar(10) COLLATE "default",
"bx_url" varchar(255) COLLATE "default",
"status" int2,
"bx_name" varchar(500) COLLATE "default",
"bx_image" varchar(255) COLLATE "default",
"bx_desc" text COLLATE "default",
"bx_directions" text COLLATE "default",
CONSTRAINT "venue_pkey" PRIMARY KEY ("id")
)
WITH (OIDS=FALSE)
;`

const dropVenueSql = "drop table venue"
