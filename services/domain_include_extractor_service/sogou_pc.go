package domain_include_extractor_service

import (
	"gitlab.fxt.cn/fxt/rank-util/structs/include_extractor"
	"regexp"
	"strconv"
	"strings"
)

func DomainIncludeExtractorSogouPc(req *include_extractor.IncludeExtractorRequest) (*include_extractor.DomainIncludeExtractorResponse, error){
	includeExtractorResponse := &include_extractor.DomainIncludeExtractorResponse{}
	re := regexp.MustCompile(`找到约(.*?)条结果`)
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
