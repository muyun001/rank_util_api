package rank_extractor_service

import (
	"encoding/json"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"strings"
)

type CrawlResultSmMobile struct {
	Q string `json:"q"`
	R []R
}

type R struct {
	W string `json:"w"`
}

func RankExtractorSugSmMobile(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	data := &CrawlResultSmMobile{}
	if err := json.Unmarshal([]byte(req.Body), &data); err != nil {
		return nil, err
	}

	rank := -2
	var ranks []int
	for i, item := range data.R {
		target := strings.ToLower(item.W)
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
