package keyword_include_extractor_service_test

import (
	"io/ioutil"
	"gitlab.fxt.cn/fxt/rank-util/services/keyword_include_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/include_extractor"
	"strings"
	"testing"
)

const dataFile360Pc = "./test_html/so_pc.html"

func TestIncludeExtractor360Pc(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFile360Pc)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	req := &include_extractor.IncludeExtractorRequest{
		Body:   html,
	}

	res, err := keyword_include_extractor_service.KeywordIncludeExtractor360Pc(req)
	if err != nil {
		t.Fatal("解析失败")
	}

	expectIsIncluded := false
	if res.IsIncluded != expectIsIncluded {
		t.Fatalf("期待收录情况%v,实际解出%v", expectIsIncluded, res.IsIncluded)
	}
}
