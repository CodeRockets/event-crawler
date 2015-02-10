package foursquare

type ExploreRes struct {
	Response struct {
		Groups []struct {
			Items []struct {
				MVenue FsVenue `json:"venue"`
			} `json:"items"`
		} `json:"groups"`
	} `json:"response"`
}
type GetRes struct {
	Response struct {
		MVenue FsVenue `json:"venue"`
	} `json:"response"`
}

type FsVenue struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Contact struct {
		Phone            string `json:"phone"`
		FormattedPhone   string `json:"formattedPhone"`
		Twitter          string `json:"twitter"`
		Facebook         string `json:"facebook"`
		FacebookUsername string `json:"facebookUsername"`
		FacebookName     string `json:"facebookName"`
	} `json:"contact"`
	Location struct {
		Address          string   `json:"address"`
		Lat              float32  `json:"lat"`
		Lng              float32  `json:"lng"`
		City             string   `json:"city"`
		Country          string   `json:"country"`
		FormattedAddress []string `json:"formattedAddress"`
	} `json:"location"`

	Url           string  `json:"url"`
	Rating        float32 `json:"rating"`
	RatingSignals int     `json:"ratingSignals"`
	StoreId       string  `json:"storeId"`
}
