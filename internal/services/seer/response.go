package seer

type Response struct {
	Results []Item `json:"results"`
}

type Item struct {
	Job        string  `json:"job"`
	Confidence float64 `json:"confidence"`
}
