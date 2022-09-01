package seer

type Response struct {
	Item []Item `json:"each_job"`
}

type Item struct {
	Job        string  `json:"job"`
	Confidence float64 `json:"confidence"`
}
