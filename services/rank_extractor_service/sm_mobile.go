package rank_extractor_service

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"strings"
)

func RankExtractorSmMobile(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	if false == strings.Contains(req.Body, "</html>") {
		return nil, errors.New("不是有效的html页面")
	}

	if strings.Contains(req.Body, `id="no-result-info card"`) ||
		strings.Contains(req.Body, "抱歉，没有找到") {
		return &rank_extractor.RankExtractorResponse{}, nil
	}

	rank := req.StartRank
	var ranks []int
	ckDomain := utils.TopDomain(req.CheckMatch)
	if ckDomain == "" {
		return nil, errors.New("CheckMatch域名无效")
	}

	dom.Find("div#results>div").Each(func(i int, selection *goquery.Selection) {
		rank += 1
		adUrl := selection.AttrOr("ad_dot_url", "")
		if adUrl != "" {
			return
		}

		indexedUrl := selection.Find("a").AttrOr("href", "")
		if indexedUrl == "" {
			link := selection.Find("a.c-header-inner").AttrOr("href", "")
			if link == "" {
				link := selection.Find("div.c-container>a.c-nature--v1_0_0").AttrOr("href", "")
				if link == "" {
					link := selection.Find("a.wemedia_header").AttrOr("href", "")
					if link == "" {
						return
					}
				}
			}
		}
		uncurl, err := url.ParseRequestURI(utils.FormatUrl(indexedUrl))
		if err != nil {
			return
		}
		showDomain := utils.TopDomain(uncurl.Host)
		indexedUrl2 := selection.Find("a").AttrOr("data-recoorgi", "")
		if indexedUrl2 == "" {
			indexedUrl2 = selection.Find("a.c-header-inner").AttrOr("data-recoorgi", "")
			if indexedUrl2 == "" {
				indexedUrl2 = selection.Find("div.c-container>a.c-nature--v1_0_0").AttrOr("data-recoorgi", "")
				if indexedUrl2 == "" {
					indexedUrl2 = selection.Find("a.wemedia_header").AttrOr("data-recoorgi", "")
					if indexedUrl2 == "" {
						return
					}
				}
			}
		}

		parsedURL, err := url.ParseRequestURI(utils.FormatUrl(indexedUrl2))
		if err != nil {
			return
		}
		showDomain2 := utils.TopDomain(parsedURL.Host)
		if strings.Compare(showDomain, ckDomain) == 0 || strings.Compare(showDomain2, ckDomain) == 0 {
			ranks = append(ranks, rank)
		}
	})
	rankExtractorResponse := &rank_extractor.RankExtractorResponse{}
	rankExtractorResponse.Ranks = ranks

	return rankExtractorResponse, nil
}
