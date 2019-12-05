package rank_extractor_service_test

import (
	"io/ioutil"
	"gitlab.fxt.cn/fxt/rank-util/services/rank_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"testing"
)

const dataFileSugBaiduPc = "./test_html/sug_baidu_pc.html"

func TestRankExtractorSugBaiduPc(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileSugBaiduPc)

	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := string(contents)
	req := &rank_extractor.RankExtractorRequest{
		StartRank:  0,
		Body:       html,
		CheckMatch: "油漆怎么洗",
		SiteName:   "",
	}
	res, err := rank_extractor_service.RankExtractorSugBaiduPc(req)

	if err != nil {
		t.Fatalf("解析出错: %s", err.Error())
	}

	if len(res.Ranks) == 0 {
		t.Fatal("未能检测到排名")
	}

	expectRank := 2

	if res.Ranks[0] != expectRank {
		t.Fatalf("期待排名%d,实际解出%d", expectRank, res.Ranks[0])
	}
}
