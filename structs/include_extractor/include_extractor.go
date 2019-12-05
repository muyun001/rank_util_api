package include_extractor

type IncludeExtractorRequest struct {
	Body string `json:"body" bson:"body"`
}

type DomainIncludeExtractorResponse struct {
	IncludeNum int `json:"include_num"`
}

type KeywordIncludeExtractorResponse struct {
	IsIncluded bool `json:"is_included"`
}
