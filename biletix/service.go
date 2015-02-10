package biletix

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	c "github.com/event-crawler/config"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"time"
)

func CrawlVenue(link string) (BxVenue, error) {
	client := getHttpClient()
	res, err := client.Get(link)

	if err != nil {
		return BxVenue{}, err
	}
	defer res.Body.Close()
	contents, _ := ioutil.ReadAll(res.Body)

	buf := bytes.NewBuffer(contents)
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return BxVenue{}, err
	}
	bxVenue := venueMapper(link, doc)

	return bxVenue, err
}

func CrawlEvent(link string) (BxEvent, error) {
	client := getHttpClient()
	res, err := client.Get(link)
	if err != nil {
		return BxEvent{}, err
	}
	defer res.Body.Close()
	contents, _ := ioutil.ReadAll(res.Body)

	buf := bytes.NewBuffer(contents)
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return BxEvent{}, err
	}
	bxEvent := eventMapper(link, doc)

	return bxEvent, err

}

func getHttpClient() *http.Client {
	cookieJar, _ := cookiejar.New(nil)

	cookie := http.Cookie{
		Name:    *c.BxIdCookieName,
		Value:   *c.BxIdCookie,
		Domain:  *c.BxDomainRoot,
		Expires: time.Now().AddDate(100, 0, 0),
	}
	cookies := []*http.Cookie{&cookie}

	u, _ := url.Parse(*c.BxDomain)
	cookieJar.SetCookies(u, cookies)
	return &http.Client{
		Jar: cookieJar,
	}

}
