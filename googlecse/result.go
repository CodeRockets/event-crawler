package googlecse

type CseResult struct {
	Items []struct {
		Link string `json:"link"`
	} `json:"items"`
}
