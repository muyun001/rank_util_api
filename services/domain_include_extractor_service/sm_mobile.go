package domain_include_extractor_service

import (
	"gitlab.fxt.cn/fxt/rank-util/structs/include_extractor"
	"regexp"
	"strconv"
	"strings"
)

func DomainIncludeExtractorSmMobile(req *include_extractor.IncludeExtractorRequest) (*include_extractor.DomainIncludeExtractorResponse, error){
	includeExtractorResponse := &include_extractor.DomainIncludeExtractorResponse{}
	re := regexp.MustCompile(`神马收录该网站约<i>(.*?)</i>个`)
	subMatch := re.FindStringSubmatch(req.Body)
	if len(subMatch) == 2 {
		includeNum, err := strconv.Atoi(strings.Replace(subMatch[1], ",", "", -1))
		if err != nil {
			return includeExtractorResponse, err
		}
		includeExtractorResponse.IncludeNum = includeNum
	}

	return includeExtractorResponse, nil
}
