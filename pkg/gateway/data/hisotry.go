package data

// GET "https://atcoder.jp/users/" + user + "/history/json"
type History struct {
	IsRated           bool   `json:"IsRated"`
	Place             int32  `json:"Place"`
	OldRating         int32  `json:"OldRating"`
	NewRating         int32  `json:"NewRating"`
	Performance       int32  `json:"Performance"`
	InnerPerformance  int32  `json:"InnerPerformance"`
	ContestScreenName string `json:"ContestScreenName"`
	ContestName       string `json:"ContestName"`
	ContestNameEn     string `json:"ContestNameEn"`
	EndTime           string `json:"EndTime"`
}
