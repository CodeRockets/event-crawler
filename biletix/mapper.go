package biletix

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
	"time"
)

func venueMapper(link string, doc *goquery.Document) BxVenue {

	img, _ := doc.Find(".venueimage img").Attr("src")
	bxVenue := BxVenue{
		Name:       doc.Find(".venuename h1").Text(),
		Image:      img,
		Desc:       doc.Find(".innerslide p").Text(),
		Directions: doc.Find("#tab3 .epcontent p").Text(),
		UpdateDate: time.Now(),
		Status:     3,
		Tags:       strings.Replace(strings.Replace(strings.Replace(link, "/ISTANBUL/tr", "", -1), "http://www.biletix.com/mekan/", "", -1), "/TURKIYE/tr", "", -1)}

	doc.Find("#venueresultlists .grid_14").Each(func(i int, s *goquery.Selection) {
		eventLink, _ := s.Find(".grid_6 .EventNameLink").Attr("href")
		if eventLink != "" {
			bxVenue.BxEventLinks = append(bxVenue.BxEventLinks, "http://www.biletix.com"+eventLink)
		}
	})
	return bxVenue

}
func eventMapper(link string, doc *goquery.Document) BxEvent {
	html, _ := doc.Html()

	//Biletix Id
	r, _ := regexp.Compile("primary_tag((.)*)\"")
	bxId := strings.Replace(strings.Replace(r.FindString(string(html)), "primary_tag", "", -1), "\"", "", -1)
	//Biletix Link
	r, _ = regexp.Compile("detail:((.)*)\"")
	bxLink := strings.Replace(strings.Replace(r.FindString(string(html)), "detail:", "", -1), "\"", "", -1)
	//Start Date
	r, _ = regexp.Compile("start:((.|\n)*)},")
	datehtml := r.FindString(string(html))
	r, _ = regexp.Compile("date:((.|)*)\",")
	date := strings.Replace(strings.Replace(strings.Replace(r.FindString(string(datehtml)), "date:", "", -1), "\"", "", -1), ",", "", -1)
	r, _ = regexp.Compile("time:((.|)*)\",")
	ttime := strings.Replace(strings.Replace(strings.Replace(r.FindString(string(datehtml)), "time:", "", -1), "\"", "", -1), ",", "", -1)
	startDate, _ := time.Parse("2006-01-02 15:04", strings.TrimSpace(date)+" "+strings.TrimSpace(ttime))
	//Event Image
	r, _ = regexp.Compile("primary:((.|)*)\"")
	eventImage := strings.Replace(strings.Replace(r.FindString(string(html)), "primary:", "", -1), "\"", "", -1)
	//Purchase Link
	r, _ = regexp.Compile("purchase:((.|)*)\"")
	purchaseLink := strings.Replace(strings.Replace(r.FindString(string(html)), "purchase:", "", -1), "\"", "", -1)

	//Event Price
	var eventprice string
	eventprice = ""
	doc.Find("#eventpage_topright #tab1").Find("[itemprop]").Each(func(i int, s *goquery.Selection) {
		itempropname, exists := s.Attr("itemprop")
		if !exists {
			return
		}
		if itempropname == "price" {
			price := strings.TrimSpace(s.Text())
			if price != "" {
				eventprice += strings.TrimSpace(s.Text()) + ","
			}
		}
	})
	eventprice = "[" + strings.TrimSuffix(eventprice, ",") + "]"
	eventName := doc.Find(".eventname a #eventnameh1 span").Text()

	desc := doc.Find(".innerslide p").Text()

	return BxEvent{
		EventDesc:      desc,
		BxId:           bxId,
		BxLink:         bxLink,
		EventDateStart: startDate,
		EventImage:     eventImage,
		PurchaseLink:   purchaseLink,
		EventPrice:     eventprice,
		Name:           eventName,
	}
}
