package foursquare

import (
	"database/sql"

	"time"
)

//Inserting venue if not exists
func InsVenue(db *sql.DB, venue FsVenue) (int, error) {

	var rowCount int
	var err error
	var lastid = 0

	err = db.QueryRow("select count(*) from venue where fs_id=$1", venue.Id).Scan(&rowCount)
	if err != nil {
		return lastid, err
	}
	if rowCount == 0 {
		lastid, err = insert(venue, db)
	}
	return lastid, err
}

//Deleting venue by fs_id
func DelByFsId(db *sql.DB, id string) (sql.Result, error) {
	return db.Exec("delete from venue where fs_id=$1", id)

}

func insert(venue FsVenue, db *sql.DB) (int, error) {

	var lastId int
	err := db.QueryRow("insert into venue (id,fs_id,name,phone,twitter,facebook_id,facebook_username,lat,lon,url,city,country,address,formatted_address,fs_tags,fs_rating,fs_rating_signals,created_at,updated_at,status) VALUES (DEFAULT,$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19) RETURNING id;",
		venue.Id,
		venue.Name,
		venue.Contact.Phone,
		venue.Contact.Twitter,
		venue.Contact.Facebook,
		venue.Contact.FacebookUsername,
		venue.Location.Lat,
		venue.Location.Lng,
		venue.Url,
		venue.Location.City,
		venue.Location.Country,
		venue.Location.Address,
		venue.Location.FormattedAddress[0],
		"{'konser', 'konser salonu'}",
		venue.Rating,
		venue.RatingSignals,
		time.Now(),
		time.Now(),
		1).Scan(&lastId)
	return lastId, err

}
