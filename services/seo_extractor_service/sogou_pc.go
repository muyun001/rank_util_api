package rank_extractor_service

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"strings"
)

func RankExtractorSogouPc(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	if false == strings.Contains(req.Body, "</html>") {
		return nil, errors.New("不是有效的html页面")
	}
	if strings.Contains(req.Body, `class="icon_noRes"`) ||
		strings.Contains(req.Body, " 抱歉，没有找到与") {
		return &rank_extractor.RankExtractorResponse{}, nil
	}

	ckDomain := utils.TopDomain(req.CheckMatch)
	if ckDomain == "" {
		return nil, errors.New("CheckMatch域名无效")
	}
	rank := req.StartRank
	var ranks []int
	dom.Find("div.results>div").Each(func(i int, selection *goquery.Selection) {
		rank += 1
		indexedUrl := selection.Find("div.fb>a").AttrOr("href", "")
		if indexedUrl == "" {
			indexedUrl = selection.Find("div>h3>a").AttrOr("href", "")
		}
		if indexedUrl == "" {
			return
		}
		indexedUrl = replaceZw(indexedUrl)
		showUrl, err := url.ParseRequestURI(utils.FormatUrl(indexedUrl))
		if err != nil {
			return
		}
		showDomain := utils.TopDomain(showUrl.Host)

		if strings.Compare(showDomain, ckDomain) == 0 {
			ranks = append(ranks, rank)
		}
	})
	rankExtractorResponse := &rank_extractor.RankExtractorResponse{}
	rankExtractorResponse.Ranks = ranks

	return rankExtractorResponse, nil
}

func replaceZw(content string) string {
	listContent := strings.Split(content, "url=")
	if len(listContent) == 2 {
		content, err := url.QueryUnescape(strings.Replace(listContent[1], " ", "", -1))
		if err != nil {
			return ""
		}
		return content
	}
	return content
}
