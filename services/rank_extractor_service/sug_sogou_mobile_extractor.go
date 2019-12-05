package rank_extractor_service

import (
	"encoding/json"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"strings"
)

type CrawlResultSogouMobile struct {
	P       bool   `json:"p"`
	Q       string `json:"q"`
	S       []string
	Answer  []string
	D       []string
	AType   []string
	SuguuId string `json:"suguuid"`
	AVrId   []string
	Type    []string
	Vr      []string
}

func RankExtractorSugSogouMobile(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	req.Body = strings.Replace(req.Body, "window.sogou.sug(", "", -1)
	req.Body = strings.Replace(req.Body, ")", "", -1)

	data := &CrawlResultSogouMobile{}
	if err := json.Unmarshal([]byte(req.Body), &data); err != nil {
		return nil, err
	}

	rank := -2
	var ranks []int
	for i, item := range data.S {
		target := strings.ToLower(item)
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
