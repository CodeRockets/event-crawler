package facebook

import (
	"database/sql"
	"encoding/json"
	c "github.com/event-crawler/config"
	"io/ioutil"
	"net/http"
)

type Event struct {
	Id        int
	FbId      string `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	Cover     struct {
		Source string `json:"source"`
	} `json:"cover"`
	Description string `json:"description"`
	Ticket_Url  string `json:"ticket_uri"`
	Attending   int    `json:"attending_count"`
	Place       struct {
		VenueId string `json:"id"`
	} `json:"place"`
	PageId string
}

func (e *Event) Call() error {
	gUrl := *c.FbApiEndpoint + e.FbId + "?fields=description,name,cover,ticket_uri,attending_count,place&access_token=" + *c.FbApiAccessToken

	client := &http.Client{}
	req, _ := http.NewRequest("GET", gUrl, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	err = json.Unmarshal(res, &e)
	return err
}

func (e *Event) CheckVenueIsInDb(db *sql.DB) error {

	var rowCount int
	var err error

	err = db.QueryRow("select count(*) from pages where fb_id=$1", e.Place.VenueId).Scan(&rowCount)
	if err != nil {
		return err
	}
	if rowCount != 0 {
		return nil
	}

	page := new(Page)
	page.FbId = e.Place.VenueId
	err = page.Call()
	_, err = page.Insert(db)
	return err

}

func (e Event) Insert(db *sql.DB) (int, error) {
	var rowCount int
	err := db.QueryRow("select count(*) from events where fb_id=$1", e.FbId).Scan(&rowCount)
	if err != nil {
		return 0, err
	}
	if rowCount != 0 {
		return rowCount, nil
	}

	var lastId int
	err = db.QueryRow("INSERT INTO events(id, fb_id, name, start_time, image, description, ticket_uri, page_id,attending_count,venue_id)VALUES (DEFAULT,$1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id;",
		e.FbId,
		e.Name,
		e.StartTime,
		e.Cover.Source,
		e.Description,
		e.Ticket_Url,
		e.PageId,
		e.Attending,
		e.Place.VenueId,
	).Scan(&lastId)
	return lastId, err
}
