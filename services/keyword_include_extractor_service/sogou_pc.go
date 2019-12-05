package keyword_include_extractor_service

import (
	"github.com/PuerkitoBio/goquery"
	"gitlab.fxt.cn/fxt/rank-util/structs/include_extractor"
	"regexp"
	"strings"
)

func KeywordIncludeExtractorSogouPc(req *include_extractor.IncludeExtractorRequest) (*include_extractor.KeywordIncludeExtractorResponse, error) {
	includeExtractorResponse := &include_extractor.KeywordIncludeExtractorResponse{}
	re := regexp.MustCompile(`站内没有找到能和.*?匹配的内容。`)
	subMatch := re.FindStringSubmatch(req.Body)
	if len(subMatch) != 0 {
		return includeExtractorResponse, nil
	}

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	dom.Find("div.results>div h3.pt a").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		if i == 0 {
			selectionHtml, err := selection.Html()
			if err != nil {
				return false
			}

			re := regexp.MustCompile(`<em>.*?</em>`)
			subMatch := re.FindStringSubmatch(selectionHtml)
			if len(subMatch) != 0 {
				includeExtractorResponse.IsIncluded = true
			}
		}
		return false
	})

	return includeExtractorResponse, nil
}
