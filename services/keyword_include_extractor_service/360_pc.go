package keyword_include_extractor_service

import (
	"gitlab.fxt.cn/fxt/rank-util/structs/include_extractor"
	"regexp"
)

func KeywordIncludeExtractor360Pc(req *include_extractor.IncludeExtractorRequest) (*include_extractor.KeywordIncludeExtractorResponse, error){
	includeExtractorResponse := &include_extractor.KeywordIncludeExtractorResponse{}
	re := regexp.MustCompile(`抱歉，未找到和.*?相关的网页`)
	subMatch := re.FindStringSubmatch(req.Body)
	if len(subMatch) == 0 {
		includeExtractorResponse.IsIncluded = true
		return includeExtractorResponse, nil
	}

	return includeExtractorResponse, nil
}
