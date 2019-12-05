package keyword_include_extractor_service

import (
	"gitlab.fxt.cn/fxt/rank-util/structs/include_extractor"
	"regexp"
)

func KeywordIncludeExtractorBaiduPc(req *include_extractor.IncludeExtractorRequest) (*include_extractor.KeywordIncludeExtractorResponse, error) {
	includeExtractorResponse := &include_extractor.KeywordIncludeExtractorResponse{}
	re := regexp.MustCompile(`很抱歉，没有找到与.*?相关的网页`)
	subMatch := re.FindStringSubmatch(req.Body)
	if len(subMatch) == 0 {
		includeExtractorResponse.IsIncluded = true
		return includeExtractorResponse, nil
	}

	return includeExtractorResponse, nil
}
