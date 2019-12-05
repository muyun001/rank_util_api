package rank_extractor_service

import (
	"encoding/json"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"strings"
)

type CrawlResultBaiduPc struct {
	Q string `json:"q"`
	P bool   `json:"p"`
	G []GBaiduPc
}

type GBaiduPc struct {
	Type string `json:"type"`
	Sa   string `json:"sa"`
	Q    string `json:"q"`
}

func RankExtractorSugBaiduPc(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	data := &CrawlResultBaiduPc{}
	if err := json.Unmarshal([]byte(req.Body), &data); err != nil {
		return nil, err
	}

	rank := -2
	var ranks []int
	for i, item := range data.G {
		target := strings.ToLower(item.Q)
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
