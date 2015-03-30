package facebook

import (
	"database/sql"
	"encoding/json"
	"fmt"
	c "github.com/event-crawler/config"
	"io/ioutil"
	"net/http"
	/*	"net/url" /*	"strings"*/)

type Page struct {
	Id       int
	FbId     string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Checkins int    `json:"checkins"`
	Cover    struct {
		Source string `json:"source"`
	} `json:"cover"`
	Description string `json:"description"`
	Likes       int    `json:"likes"`
	Location    struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
		Address   string  `json:"street"`
		City      string  `json:"city"`
	} `json:"location"`
	UserName string `json:"username"`
	IsVenue  bool
}

func (p Page) Insert(db *sql.DB) (int, error) {

	var rowCount int
	err := db.QueryRow("select count(*) from pages where fb_id=$1", p.FbId).Scan(&rowCount)
	if err != nil {
		return 0, err
	}
	if rowCount != 0 {
		return rowCount, nil
	}

	var lastId int
	err = db.QueryRow("INSERT INTO pages(id, fb_id, name, category, checkins, image, description, likes, latitude, longitude, street, city, username, is_venue)VALUES (DEFAULT,$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING id;",
		p.FbId,
		p.Name,
		p.Category,
		p.Checkins,
		p.Cover.Source,
		p.Description,
		p.Likes,
		p.Location.Latitude,
		p.Location.Longitude,
		p.Location.Address,
		p.Location.City,
		p.UserName,
		p.IsVenue,
	).Scan(&lastId)
	return lastId, err
}

func (p Page) Delete(db *sql.DB, id string) (sql.Result, error) {
	return db.Exec("delete from pages where fb_id=$1", id)
}

func (p *Page) Call() error {
	gUrl := *c.FbApiEndpoint + p.FbId + "?access_token=" + *c.FbApiAccessToken

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
	err = json.Unmarshal(res, &p)
	if err != nil {
		return err
	}
	return err
}

func (p *Page) GetEvents() ([]Event, error) {
	gUrl := *c.FbApiEndpoint + p.FbId + "/events?access_token=" + *c.FbApiAccessToken
	fmt.Println(gUrl)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", gUrl, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(res, &objmap)

	var eventlist []Event

	err = json.Unmarshal(*objmap["data"], &eventlist)
	if err != nil {
		return nil, err
	}
	return eventlist, err
}
