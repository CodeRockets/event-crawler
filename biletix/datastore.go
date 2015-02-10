package biletix

import (
	"database/sql"
	"time"
)

//Inserting venue if not exists
func UpdateVenue(db *sql.DB, bxvenue BxVenue) (sql.Result, error) {
	return db.Exec("UPDATE venue SET bx_name = $1, bx_image = $2, bx_desc = $3,bx_directions = $4,updated_at=$5, status=$6 , bx_tag=$7 WHERE id = $8 ",
		bxvenue.Name,
		bxvenue.Image,
		bxvenue.Desc,
		bxvenue.Directions,
		bxvenue.UpdateDate,
		bxvenue.Status,
		bxvenue.Tags,
		bxvenue.VenueId)
}

func UpdateVenueBxUrlByFsId(db *sql.DB, fsId string, bxUrl string) (sql.Result, error) {
	return db.Exec("UPDATE venue SET bx_url=$1 WHERE fs_id=$2", bxUrl, fsId)
}

func InsertEvent(db *sql.DB, bxevent BxEvent) (sql.Result, error) {

	return db.Exec("INSERT INTO event (venue_id,event_name,bx_event_link,created_at,updated_at,status,event_image,bx_event_id,event_desc,event_price,event_date_start,purchase_link) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)",
		bxevent.VenueId,
		bxevent.Name,
		bxevent.BxLink,
		time.Now(),
		time.Now(),
		2,
		bxevent.EventImage,
		bxevent.BxId,
		bxevent.EventDesc,
		bxevent.EventPrice,
		bxevent.EventDateStart,
		bxevent.PurchaseLink)
}

//Deleting venue by fs_id
func DelEventByBxId(db *sql.DB, bxid string) (sql.Result, error) {
	return db.Exec("DELETE FROM event WHERE bx_event_id=$1", bxid)

}
