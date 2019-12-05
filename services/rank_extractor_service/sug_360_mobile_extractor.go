package rank_extractor_service

import (
	"encoding/json"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"strings"
)

type CrawlData360Mobile struct {
	Errno int `json:"errno"`
	Data  Data
	Msg   string `json:"msg"`
}

type Data struct {
	Query   string `json:"query"`
	Sug     []Sug
	Version string `json:"version"`
}

type Sug struct {
	Word string `json:"word"`
}

func RankExtractorSug360Mobile(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	data := &CrawlData360Mobile{}
	if err := json.Unmarshal([]byte(req.Body), &data); err != nil {
		return nil, err
	}

	rank := -2
	var ranks []int
	for i, item := range data.Data.Sug {
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
