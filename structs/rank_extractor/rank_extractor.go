package rank_extractor

type RankExtractorRequest struct {
	StartRank  int    `json:"start_rank" bson:"start_rank"`
	Body       string `json:"body" bson:"body"`
	CheckMatch string `json:"check_match" bson:"check_match"`
	SiteName   string `json:"site_name" bson:"site_name"`
}

type RankExtractorResponse struct {
	Ranks []int `json:"ranks"`
}
