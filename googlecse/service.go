package googlecse

import (
	"encoding/json"

	c "github.com/event-crawler/config"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SearchBxLinks(keyword string) (*CseResult, error) {

	var Url *url.URL
	Url, _ = url.Parse(*c.GoogleCSEEndpoint)
	parameters := url.Values{}
	parameters.Add("q", keyword)
	parameters.Add("key", *c.GoogleCSEApiKey)
	parameters.Add("cx", *c.GoogleCSEAppCsx)
	Url.RawQuery = parameters.Encode()

	resp_body_search, _ := getRequest(Url.String())
	var gcr CseResult
	err := json.Unmarshal(resp_body_search, &gcr)
	return &gcr, err
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
