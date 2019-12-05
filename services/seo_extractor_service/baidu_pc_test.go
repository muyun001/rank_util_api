package rank_extractor_service_test

import (
	"gitlab.fxt.cn/fxt/rank-util/services/rank_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"io/ioutil"
	"strings"
	"testing"
)

const dataFileBaiduPc = "./test_html/baidu_pc.html"
const dataFileBaiduPcDot = "./test_html/baidu_pc_dot.html"
const dataFileBaiduPcSiteName = "./test_html/baidu_pc_site_name.html"

func TestRankExtractorBaiduPc(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileBaiduPc)

	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	req := &rank_extractor.RankExtractorRequest{
		StartRank:  0,
		Body:       html,
		CheckMatch: "www.cq1987.com",
		SiteName:   "重庆勇者教育科技公司",
	}
	res, err := rank_extractor_service.RankExtractorBaiduPc(req)

	if err != nil {
		t.Fatalf("解析出错: %s", err.Error())
	}

	if len(res.Ranks) == 0 {
		t.Fatal("未能检测到排名")
	}

	expectRank := 1

	if res.Ranks[0] != expectRank {
		t.Fatalf("期待排名%d,实际解出%d", expectRank, res.Ranks[0])
	}
}

func TestRankExtractorBaiduPc2(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileBaiduPcDot)

	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	req := &rank_extractor.RankExtractorRequest{
		StartRank:  0,
		Body:       html,
		CheckMatch: "www.jinshizhongqing.com",
		SiteName:   "",
	}
	res, err := rank_extractor_service.RankExtractorBaiduPc(req)

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

func TestRankExtractorBaiduPcSiteName(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileBaiduPcSiteName)

	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	req := &rank_extractor.RankExtractorRequest{
		StartRank:  0,
		Body:       html,
		CheckMatch: "www.zhitouweilai.com",
		SiteName:   "智投未来",
	}
	res, err := rank_extractor_service.RankExtractorBaiduPc(req)

	if err != nil {
		t.Fatalf("解析出错: %s", err.Error())
	}

	if len(res.Ranks) == 0 {
		t.Fatal("未能检测到排名")
	}

	if len(res.Ranks) != 2 {
		t.Fatalf("期待排名数量2,实际解出%d", len(res.Ranks))
	}

	expectRanks := []int{3, 5}

	if res.Ranks[0] != expectRanks[0] || res.Ranks[1] != expectRanks[1] {
		t.Fatalf("期待排名%d,%d,实际解出%d,%d", expectRanks[0], expectRanks[1], res.Ranks[0], res.Ranks[1])
	}
}
