package domain_include_extractor_service

import (
	"gitlab.fxt.cn/fxt/rank-util/structs/include_extractor"
	"regexp"
	"strconv"
	"strings"
)

func DomainIncludeExtractor360Pc(req *include_extractor.IncludeExtractorRequest) (*include_extractor.DomainIncludeExtractorResponse, error){
	includeExtractorResponse := &include_extractor.DomainIncludeExtractorResponse{}
	regexpStrs := []string{`找到相关结果<strong>(.*?)</strong>个`, `该网站约(.*?)个网页被360搜索收录`}
	for _, reg := range regexpStrs {
		re := regexp.MustCompile(reg)
		subMatch := re.FindStringSubmatch(req.Body)
		if len(subMatch) == 2 {
			includeNum, err := strconv.Atoi(strings.Replace(subMatch[1], ",", "", -1))
			if err != nil {
				return includeExtractorResponse, err
			}
			includeExtractorResponse.IncludeNum = includeNum
			break
		}
	}

	return includeExtractorResponse, nil
}
