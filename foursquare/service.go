package foursquare

import (
	"encoding/json"
	/*	"fmt"*/
	c "github.com/event-crawler/config"
	"io/ioutil"
	"net/http"

	"net/url"
	"strings"
)

func Explore(near string, tag string) (*ExploreRes, error) {
	near = strings.Replace(url.QueryEscape(near), "+", "%20", -1)
	tag = strings.Replace(url.QueryEscape(tag), "+", "%20", -1)
	eUrl := *c.FsExploreEndpoint + "&near=" + near + "&query=" + tag + "&client_id=" + *c.FsClientId + "&client_secret=" + *c.FsClientSecret + "&v=20141028"

	res, err := getRequest(eUrl)
	if err != nil {
		return nil, err
	}
	var fvrExplore ExploreRes
	err1 := json.Unmarshal(res, &fvrExplore)
	return &fvrExplore, err1
}

func Get(id string) (*GetRes, error) {
	gUrl := *c.FsGetEndpoint + id + "?client_id=" + *c.FsClientId + "&client_secret=" + *c.FsClientSecret + "&v=20141028"
	res, err := getRequest(gUrl)
	if err != nil {
		return nil, err
	}
	var fvrGet GetRes
	err = json.Unmarshal(res, &fvrGet)
	return &fvrGet, err
}

func getRequest(url string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	return res, err
}
