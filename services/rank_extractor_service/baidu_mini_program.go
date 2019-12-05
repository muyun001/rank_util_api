package rank_extractor_service

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"regexp"
	"strings"
)

func RankExtractorBaiduMiniProgram(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	if false == strings.Contains(req.Body, "</html>") {
		return nil, errors.New("不是有效的html页面")
	}

	var ranks []int
	rank := req.StartRank
	dom.Find("div.c-result").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		if len(selection.Find("div.c-result-content ").Nodes) == 0 {
			return true
		}

		rank += 1
		dataLog := selection.AttrOr("data-log", "")
		if dataLog != "" {
			dataLog = strings.Replace(dataLog, "'", "\"", -1)

			var sxData map[string]string
			if err := json.Unmarshal([]byte(dataLog), &sxData); err != nil {
				return true
			}

			if sxData["mu"] == "" {
				return true
			}

			appKeyContainer := selection.Find("div.c-result-content article")
			if len(appKeyContainer.Nodes) == 0 {
				return true
			}

			appKeyStr := appKeyContainer.AttrOr("rl-link-data-xcx", "")
			if appKeyStr == "false" {
				return true
			}

			r := regexp.MustCompile(`"xcxAppKey":"(.*?)",`)
			subMatch := r.FindStringSubmatch(appKeyStr)
			if len(subMatch) == 2 {
				if subMatch[1] == req.CheckMatch {
					ranks = append(ranks, rank)
				}
			}
		}
		return true
	})

	rankExtractorResponse := &rank_extractor.RankExtractorResponse{}
	rankExtractorResponse.Ranks = ranks

	return rankExtractorResponse, nil
}
