package rank_extractor_service_test

import (
	"io/ioutil"
	"gitlab.fxt.cn/fxt/rank-util/services/rank_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"strings"
	"testing"
)

const dataFileBaiduMobile = "./test_html/baidu_mobile.html"
const dataFileBaiduMobile2 = "./test_html/baidu_mobile2.html"

func TestRankExtractorBaiduMobile(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileBaiduMobile)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	req := &rank_extractor.RankExtractorRequest{
		StartRank:  1,
		Body:       html,
		CheckMatch: "www.sctiesiwang.com",
		SiteName:   "万盛达",
	}
	res, err := rank_extractor_service.RankExtractorBaiduMobile(req)

	if err != nil {
		t.Fatalf("解析出错: %s", err.Error())
	}

	if len(res.Ranks) == 0 {
		t.Fatal("未能检测到排名")
	}

	expectRank := 4

	if res.Ranks[0] != expectRank {
		t.Fatalf("期待排名%d,实际解出%d", expectRank, res.Ranks[0])
	}
}

func TestRankExtractorBaiduMobile2(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileBaiduMobile2)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	req := &rank_extractor.RankExtractorRequest{
		StartRank:  20,
		Body:       html,
		CheckMatch: "www.autodoortech.com.cn",
		SiteName:   "苏州比利孚建筑安装",
	}
	res, err := rank_extractor_service.RankExtractorBaiduMobile(req)

	if err != nil {
		t.Fatalf("解析出错: %s", err.Error())
	}

	if len(res.Ranks) == 0 {
		t.Fatal("未能检测到排名")
	}

	if len(res.Ranks) != 2 {
		t.Fatalf("期待排名数量2,实际解出%d", len(res.Ranks))
	}

	expectRanks := []int{28, 29}

	if res.Ranks[0] != expectRanks[0] || res.Ranks[1] != expectRanks[1] {
		t.Fatalf("期待排名%d,%d,实际解出%d,%d", expectRanks[0], expectRanks[1], res.Ranks[0], res.Ranks[1])
	}
}
