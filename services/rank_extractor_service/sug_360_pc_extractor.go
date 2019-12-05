package rank_extractor_service

import (
	"encoding/json"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"strings"
)

type CrawlData360Pc struct {
	Ext     string `json:"ext"`
	Query   string `json:"query"`
	Tag     string `json:"tag"`
	SSid    string `json:"ssid"`
	Version string `json:"version"`
	Result  []Result360Pc
}

type Result360Pc struct {
	Rank   string `json:"rank"`
	ReSrc  string `json:"resrc"`
	Source string `json:"source"`
	Word   string `json:"word"`
}

func RankExtractorSug360Pc(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	data := &CrawlData360Pc{}
	if err := json.Unmarshal([]byte(req.Body), &data); err != nil {
		return nil, err
	}

	rank := -2
	var ranks []int
	for i, item := range data.Result {
		target := strings.ToLower(item.Word)
		check := strings.ToLower(req.CheckMatch)
		if utils.IsSugContains(check, target) {
			rank = i + 1
			ranks = append(ranks, rank)
		}
	}
	rankExtractorResponse := &rank_extractor.RankExtractorResponse{}
	rankExtractorResponse.Ranks = ranks

	return rankExtractorResponse, nil
}
