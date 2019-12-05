package rank_extractor_service_test

import (
	"io/ioutil"
	"gitlab.fxt.cn/fxt/rank-util/services/rank_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"strings"
	"testing"
)

const dataFileBaiduMiniProgram = "./test_html/baidu_mini_program.html"

func TestRankExtractorBaiduMiniProgram(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileBaiduMiniProgram)

	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	req := &rank_extractor.RankExtractorRequest{
		StartRank:  0,
		Body:       html,
		CheckMatch: "sPY6Pm6NyfWFqt30XpeSoEYlryV2G82E",
		SiteName:   "",
	}
	res, err := rank_extractor_service.RankExtractorBaiduMiniProgram(req)

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
