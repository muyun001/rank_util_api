package rank_extractor_service

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"strings"
)

func RankExtractor360Pc(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(req.Body))
	if err != nil {
		log.Fatalln(err)
	}

	if false == strings.Contains(req.Body, "</html>") {
		return nil, errors.New("不是有效的html页面")
	}

	if strings.Contains(req.Body, `id="tips"`) ||
		strings.Contains(req.Body, "抱歉，未找到和 ") {
		return &rank_extractor.RankExtractorResponse{}, nil
	}

	ckDomain := utils.TopDomain(req.CheckMatch)
	if ckDomain == "" {
		return nil, errors.New("CheckMatch域名无效")
	}

	rank := req.StartRank
	var ranks []int
	indexedUrl := ""
	dom.Find("li.res-list").Find("h3 a").Each(func(i int, selection *goquery.Selection) {
		rank += 1
		indexedUrl = selection.AttrOr("data-url", "")
		if indexedUrl == "" {
			indexedUrl = selection.AttrOr("href", "")
			if indexedUrl == "" {
				return
			}
			if strings.Index(indexedUrl, "https://www.so.com/link?") > -1 {
				indexedUrl, err = url.QueryUnescape(strings.Split(indexedUrl, "url=")[1])
				if err != nil {
					return
				}
			}
		}
		showUrl, err := url.ParseRequestURI(utils.FormatUrl(indexedUrl))
		if err != nil {
			return
		}

		showDomain := utils.TopDomain(showUrl.Host)
		if strings.Compare(ckDomain, showDomain) == 0 {
			ranks = append(ranks, rank)
		}
	})
	rankExtractorResponse := &rank_extractor.RankExtractorResponse{}
	rankExtractorResponse.Ranks = ranks

	return rankExtractorResponse, nil
}
