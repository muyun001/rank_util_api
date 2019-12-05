package rank_extractor_service

import (
	"fmt"
	"github.com/tidwall/gjson"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"strings"
)

func RankExtractorSugSogouPc(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	req.Body = strings.Replace(req.Body, "window.sogou.sug(", "", -1)
	req.Body = strings.Replace(req.Body, ",-1);", "", -1)

	rank := -2
	var ranks []int
	sugJson := gjson.Parse(fmt.Sprintf(`{"data":%s}`, req.Body))
	if sugJson.Get("data").IsArray() {
		dataArr := sugJson.Get("data").Array()
		if len(dataArr) >= 2 {
			sugsResultArry := dataArr[1].Array()
			for i, sugResult := range sugsResultArry {
				target := strings.ToLower(sugResult.Str)
				check := strings.ToLower(req.CheckMatch)
				if utils.IsSugContains(check, target) {
					rank = i + 1
					ranks = append(ranks, rank)
				}
			}
		}
	}
	rankExtractorResponse := &rank_extractor.RankExtractorResponse{}
	rankExtractorResponse.Ranks = ranks

	return rankExtractorResponse, nil
}
