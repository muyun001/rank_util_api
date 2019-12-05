package rank_extractor_service

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"

	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"strings"
)

func RankExtractorBaiduMobile(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	if false == strings.Contains(req.Body, "</html>") {
		return nil, errors.New("不是有效的html页面")
	}
	rank := req.StartRank
	var ranks []int
	ckDomain := utils.TopDomain(req.CheckMatch)
	if ckDomain == "" {
		return nil, errors.New("CheckMatch域名无效")
	}
	dom.Find("div[order].c-result").Each(func(i int, selection *goquery.Selection) {
		rank += 1
		muUri := ""
		dataLog := selection.AttrOr("data-log", "")
		if dataLog == "" {
			return
		}
		dataLog = strings.Replace(dataLog, "'", "\"", -1)

		var sxData map[string]string
		if err := json.Unmarshal([]byte(dataLog), &sxData); err != nil {
			return
		}
		muUri = sxData["mu"]
		if muUri == "" {
			return
		}
		muURL, err := url.ParseRequestURI(utils.FormatUrl(muUri))
		if err != nil {
			return
		}
		muDomain := utils.TopDomain(muURL.Host)
		if strings.Compare(muDomain, ckDomain) == 0 {
			ranks = append(ranks, rank)
		}
	})
	rankExtractorResponse := &rank_extractor.RankExtractorResponse{}
	rankExtractorResponse.Ranks = ranks

	return rankExtractorResponse, nil
}
