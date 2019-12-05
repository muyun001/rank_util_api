package keyword_include_extractor_service

import (
	"github.com/PuerkitoBio/goquery"
	"gitlab.fxt.cn/fxt/rank-util/structs/include_extractor"
	"regexp"
	"strings"
)

func KeywordIncludeExtractorSmMobile(req *include_extractor.IncludeExtractorRequest) (*include_extractor.KeywordIncludeExtractorResponse, error){
	includeExtractorResponse := &include_extractor.KeywordIncludeExtractorResponse{}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	dom.Find("div#results>div span[c-bind]").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		if i == 0{
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